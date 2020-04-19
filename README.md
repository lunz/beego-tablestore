# Alibaba Cloud TableStore & Beego API

When I shop around the cheap and relaible no-sql products in the market, Alibaba TableStore caught my eye. The free trial period spans 12 months!! In addition, each year, Alibaba cloud is capable to handle the giant volume spike on double 11. There must be a strong cloud platform behind to make this happening. Let me try out. 

TableStore is a competing NO-SQL product in the market. I was impressed by the pagination feature TableStore provided. The next page token is the primary key set which can be customized in fly. Usually the pagination of No-SQL databases is not friendly and cannot be changed once starts. When I worked on a NO-SQL database, the generated next page token could reach up to 12KB. In addition, I can only pull page by page and even cannot change the page size. While Alibaba TableStore solved this problem easily. 

Start a request with the primary keys of the next page 1st record, along with the page size, the response will include all qualifed rows and the primary keys of the next page 1st record. The parimary keys are similar to the primary keys defined in the relationonal db such as mySql. Since primary key values are inserted by customer, primary keys and page size could be changed in fly.

A classic scenario is a product page with tons of customer comments. Usually the UX either pre-fetches a huge amount of comments or caches/accumulates the paged result per reuqest in browser to allow customer to switch page, which requires the UX to come up the complex logic to handle the cache. With Tablestore, UX only needs to cache the next page primary key sets. When customer clicks any page, use that primary key sets to get rows from db directly. Such complex logic is no longer needed. 

This example is the end-2-end API solution demonstrating creating table, adding row, deleting row, updating row, then fetching rows page by page. 

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
Usually table creation only needs once. So the table name and column definitions are hard-coded. 

  POST:    `http://:8080/v1/comment/table/create`
  
### Create or update a comment
Create or Update is using the same endpoint. 
  - If primary key sets are same, this is the update. 
  - If commentId is not supplied, this is the create.

Removing `commentId` will create a new record with an auto-generated commentId. Primary key set "prodId" and "commentId" are returned in response if this record is new.

  > POST:    `http://:8080/v1/comment`
      
      {
        "ProdId":"123",
        "CommentId": 1585518496615000,   
        "Content":"I feel greate",
      }

### Get single comment record by primary key set

  > GET:    `http://:8080/v1/comment/123/1585518496615000`
  

### Get the paged comments by prodId or prodId and commentId

  - First Page
    > GET:    `http://:8080/v1/comment/batch/123`

  - Next Page when next page primary key set is available
    > GET:    `http://:8080/v1/comment/batch/123/1585518496615000`

### Delete single comment record by primary key set

  > DELETE:    `http://:8080/v1/comment/123/1585518496615000`

## Reference
 * [Alibaba Cloud TableStore Golang SDK](https://github.com/aliyun/alibaba-cloud-sdk-go)
 * [Beego Framework](https://github.com/astaxie/beego)
 * [Set up Alibaba Free Trial Account](https://us.alibabacloud.com)
 
