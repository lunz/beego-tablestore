# Alibaba Cloud TableStore & Beego API

When I shop around the cheap and relaible no-sql products in the market, Alibaba TableStore comes into my picture. The free trial period spans 12 months!! In addition, each year, Alibaba cloud has to handle the giant colume spike on double 11. There must be a strong cloud platform behind to make this happening. Let me try out. 

TableStore is a competing NO-SQL product in the market. I was impressed by the pagination feature TableStore provided. The next page token is the primary key set. Querying page result is determined by your requested Primary Keys and Page Size. As we know, the pagination is not easy to handle for no-sql databases. The pagination feature of some No-SQL products is not friendly and hard to use. For example, when I worked on a no-sql product, the generated next page token could reach up to 12KB. While Alibaba TableStore solved this problem easily. 

Start a request with the primary keys of next page 1st record, along with the page size, the response will include all qualifed rows and the primary keys of next page 1st record. In another words, primary keys and page size could be changed randomly according to your needs without the sequence restriction. (This primary keys are similar to the PK you defined in MySql table). 

A classic scenario is the product comments. Usually the UX client either pre-fetches a huge amount of comments or caches/accumulates the paged result per reuqest to support customer switching page, which makes the client handling quite complex. Now I can get rid of those complex logic and fetch the page result upon needs. Client side only needs to cache the next page primary keys.

This example is the end-2-end solution demonstrating how to solve this classic scenario from creating table, adding row, deleting row, updating row, then fetching rows page by page. 

## Installing

* Step 1: Clone the code to your local storage. 
* Step 2: Set up a free trial account in [Alibaba Cloud](https://us.alibabacloud.com)
* Step 3: Create a tablestore instance, then get AccessKeyId, AccessKeySecret, InternetEndPoint, InstanceName
* Step 4: Open /conf/app.config, fill in AccessKeyId, AccessKeySecret, InternetEndPoint, InstanceName 
* Step 5: Open a command window, run below commands:

  ```dos
  go get ./...  # get dependencies
  go get github.com/beego/bee
  bee run
  ```
  Now the API service is up running
  

## API
### Create Table

URL:    `<url>`
Method:
Input:
```sh

```
Response:

### Post
- `Create` -

- `Update` -

### Get

### Delete

## Reference
 * [Alibaba Cloud TableStore Golang SDK](https://github.com/aliyun/alibaba-cloud-sdk-go)
 * [Beego Framework](https://github.com/astaxie/beego)
 * [Set up Alibaba Free Trial Account](https://us.alibabacloud.com)
 
