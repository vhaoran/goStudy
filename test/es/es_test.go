package es

import (
	"context"
	"fmt"
	//这里使用的是版本5，最新的是6，有改动
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/ivpusic/grpool"
	//	"gopkg.in/olivere/elastic/v7"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var host = "http://192.168.0.99:9200"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

//初始化
func init() {
	log := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(log), elastic.SetURL(host),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("######################")
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esVersion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esVersion)
	fmt.Println("---------init OK------------------")
}

func Test_es(t *testing.T) {
	create()
	delete()
	update()
	gets()

	list(10, 1)
}

func Test_mut_add(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(10, 50)
	defer pool.Release()

	// how many jobs we should wait
	max := 10000
	pool.WaitCount(max)

	// submit one or more jobs to pool
	for i := 0; i < max; i++ {
		count := i

		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()
			fmt.Println("count: ", count)

			bean := Employee{fmt.Sprint("Jane_", count), "Smith", 32, "I like to collect rock albums", []string{"music"}}
			_, err := client.Index().
				Index("test").
				//Type("employee").
				Id(fmt.Sprint(count, "_id")).
				BodyJson(bean).
				Do(context.Background())
			if err != nil {
				fmt.Println("es_test.go->", err.Error())
			} else {
				fmt.Println("ok--->", count)
			}
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()
}

func Test_get(t *testing.T) {
	gets()
}

/*下面是简单的CURD*/

//创建
func create() {
	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("megacorp").
		//Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("megacorp").
		//Type("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

//删除
func delete() {

	res, err := client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func update() {
	res, err := client.Update().
		Index("megacorp").
		//Type("employee").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

//查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("megacorp").Type("employee").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

//简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("megacorp").
		//Type("employee").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)

}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}
