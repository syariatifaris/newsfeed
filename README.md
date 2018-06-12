## Sample News Feed API

This API uses `github.com/syariatifaris/arkeus` framework

#### Database Setup

Run migration files on `files/migration`

#### Installation

1. `go get github.com/syariatifaris/newsfeed`
2. `dep ensure -v`


#### API Reference

1. GET: `news/gets/{status}` : Get all news by status, i.e: **all** (for all news), **published**, or etc.
2. POST: `news/gets/topics`: Get all by topics (OR operand), i.e: `politik`, `teknologi`

    Request:
    ```$xslt
    {
        "topics": ["teknologi"]
    }
    ```
    
    Response:
    ```$xslt
    {
        "data": [
            {
                "news_id": 13,
                "title": "Flagship terbaru dari Samsung",
                "html_content": "Samsung mengeluarkan",
                "status": "published",
                "is_valid": 1
            }
        ],
        "status": 200
    }
    ```

3. POST: `news/add`. Add news by its topic
    
    Request:
    
    ```$xslt
    {
        "news":{
            "title": "Flagship terbaru dari Samsung",
            "html_content": "Samsung mengeluarkan",
            "status": "published"
        },
        "topics": ["teknologi"]
    }
    ```
    
    Response:
    ```$xslt
    {
        "data": {
            "news": {
                "news_id": 13,
                "title": "Flagship terbaru dari Samsung",
                "html_content": "Samsung mengeluarkan",
                "status": "published"
            },
            "topics": [
                "teknologi"
            ]
        },
        "status": 200
    }
    ```
