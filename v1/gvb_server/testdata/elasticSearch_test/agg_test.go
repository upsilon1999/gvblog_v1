package elasticSearch_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)
type BucketsType struct {
	Buckets []struct {
	  KeyAsString string `json:"key_as_string"`
	  Key         int64  `json:"key"`
	  DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}
func TestDateAgg(t *testing.T) {
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	ctx := context.Background()
	// 创DateHistogram桶聚合
	aggs := elastic.NewDateHistogramAggregation().
	Field("created_at"). // 根据date字段值，对数据进行分组
	//  分组间隔：month代表每月、支持minute（每分钟）、hour（每小时）、day（每天）、week（每周）、year（每年)
	CalendarInterval("minute").
	// 设置返回结果中桶key的时间格式
	Format("yyyy-MM-dd")

	searchResult, err := client.Search().
	Index("article_index"). // 设置索引名
	Query(elastic.NewMatchAllQuery()). // 设置查询条件
	Aggregation("sales_over_time", aggs). // 设置聚合条件，并为聚合条件设置一个名字
	Size(0). // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
	Do(ctx) // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	// 使用DateHistogram函数和前面定义的聚合条件名称，查询结果
	agg, found := searchResult.Aggregations.DateHistogram("sales_over_time")
	if !found {
		logrus.Fatal("没有找到聚合数据")
	}
    if len(agg.Buckets)==0{
		logrus.Fatal("没有待展示的数据")
	}
	// 遍历桶数据
	for _, bucket := range agg.Buckets {
		// 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
		bucketValue := bucket.Key

		// 打印结果， 默认桶聚合查询，都是统计文档总数
		fmt.Printf("bucket = %v 文档总数 = %d\n", bucketValue, bucket.DocCount)
	}
}