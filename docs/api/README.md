# [Go(lang)] E-Learning #

Generated with [https://github.com/jofaval/utilities/blob/master/python/postman-to-markdown.py](https://github.com/jofaval/utilities/blob/master/python/postman-to-markdown.py)

## Contents

1. [0.auth](#0auth)
	1. [[POST] /login](#post-login)
	1. [[POST] /logout](#post-logout)
1. [Analytics](#analytics)
	1. [Dashboard](#dashboard)
		1. [[GET] /dashboards](#get-dashboards)
		1. [[GET] /dashboards/pages](#get-dashboardspages)
		1. [[GET] /dashboards by ID](#get-dashboards-by-id)
		1. [[GET] /dashboards by slug](#get-dashboards-by-slug)
		1. [[POST] /dashboards](#post-dashboards)
		1. [[PUT] /dashboards](#put-dashboards)
		1. [[DELETE] /dashboards](#delete-dashboards)
		1. [[POST] /dashboards/batch](#post-dashboardsbatch)
	1. [Query](#query)
		1. [[GET] /queries](#get-queries)
		1. [[GET] /queries/pages](#get-queriespages)
		1. [[GET] /queries by ID](#get-queries-by-id)
		1. [[POST] /queries](#post-queries)
		1. [[PUT] /queries](#put-queries)
		1. [[DELETE] /queries](#delete-queries)
		1. [[POST] /queries/batch](#post-queriesbatch)
	1. [Report](#report)
		1. [[GET] /reports](#get-reports)
		1. [[GET] /reports/pages](#get-reportspages)
		1. [[GET] /reports by ID](#get-reports-by-id)
		1. [[GET] /reports by slug](#get-reports-by-slug)
		1. [[POST] /reports](#post-reports)
		1. [[PUT] /reports](#put-reports)
		1. [[DELETE] /reports](#delete-reports)
		1. [[POST] /reports/batch](#post-reportsbatch)
1. [Courses](#courses)
	1. [dashboard](#dashboard)
		1. [[GET] /dashboards](#get-dashboards)
		1. [[GET] /dashboards/pages](#get-dashboardspages)
		1. [[GET] /dashboards by ID](#get-dashboards-by-id)
		1. [[GET] /dashboards by slug](#get-dashboards-by-slug)
		1. [[POST] /dashboards](#post-dashboards)
		1. [[PUT] /dashboards](#put-dashboards)
		1. [[DELETE] /dashboards](#delete-dashboards)
		1. [[POST] /dashboards/batch](#post-dashboardsbatch)
	1. [query](#query)
		1. [[GET] /queries](#get-queries)
		1. [[GET] /queries/pages](#get-queriespages)
		1. [[GET] /queries by ID](#get-queries-by-id)
		1. [[POST] /queries](#post-queries)
		1. [[PUT] /queries](#put-queries)
		1. [[DELETE] /queries](#delete-queries)
		1. [[POST] /queries/batch](#post-queriesbatch)
	1. [report](#report)
		1. [[GET] /reports](#get-reports)
		1. [[GET] /reports/pages](#get-reportspages)
		1. [[GET] /reports by ID](#get-reports-by-id)
		1. [[GET] /reports by slug](#get-reports-by-slug)
		1. [[POST] /reports](#post-reports)
		1. [[PUT] /reports](#put-reports)
		1. [[DELETE] /reports](#delete-reports)
		1. [[POST] /reports/batch](#post-reportsbatch)
1. [Quizzes](#quizzes)
	1. [dashboard](#dashboard)
		1. [[GET] /dashboards](#get-dashboards)
		1. [[GET] /dashboards/pages](#get-dashboardspages)
		1. [[GET] /dashboards by ID](#get-dashboards-by-id)
		1. [[GET] /dashboards by slug](#get-dashboards-by-slug)
		1. [[POST] /dashboards](#post-dashboards)
		1. [[PUT] /dashboards](#put-dashboards)
		1. [[DELETE] /dashboards](#delete-dashboards)
		1. [[POST] /dashboards/batch](#post-dashboardsbatch)
	1. [query](#query)
		1. [[GET] /queries](#get-queries)
		1. [[GET] /queries/pages](#get-queriespages)
		1. [[GET] /queries by ID](#get-queries-by-id)
		1. [[POST] /queries](#post-queries)
		1. [[PUT] /queries](#put-queries)
		1. [[DELETE] /queries](#delete-queries)
		1. [[POST] /queries/batch](#post-queriesbatch)
	1. [report](#report)
		1. [[GET] /reports](#get-reports)
		1. [[GET] /reports/pages](#get-reportspages)
		1. [[GET] /reports by ID](#get-reports-by-id)
		1. [[GET] /reports by slug](#get-reports-by-slug)
		1. [[POST] /reports](#post-reports)
		1. [[PUT] /reports](#put-reports)
		1. [[DELETE] /reports](#delete-reports)
		1. [[POST] /reports/batch](#post-reportsbatch)
1. [Roles](#roles)
	1. [dashboard](#dashboard)
		1. [[GET] /dashboards](#get-dashboards)
		1. [[GET] /dashboards/pages](#get-dashboardspages)
		1. [[GET] /dashboards by ID](#get-dashboards-by-id)
		1. [[GET] /dashboards by slug](#get-dashboards-by-slug)
		1. [[POST] /dashboards](#post-dashboards)
		1. [[PUT] /dashboards](#put-dashboards)
		1. [[DELETE] /dashboards](#delete-dashboards)
		1. [[POST] /dashboards/batch](#post-dashboardsbatch)
	1. [query](#query)
		1. [[GET] /queries](#get-queries)
		1. [[GET] /queries/pages](#get-queriespages)
		1. [[GET] /queries by ID](#get-queries-by-id)
		1. [[POST] /queries](#post-queries)
		1. [[PUT] /queries](#put-queries)
		1. [[DELETE] /queries](#delete-queries)
		1. [[POST] /queries/batch](#post-queriesbatch)
	1. [report](#report)
		1. [[GET] /reports](#get-reports)
		1. [[GET] /reports/pages](#get-reportspages)
		1. [[GET] /reports by ID](#get-reports-by-id)
		1. [[GET] /reports by slug](#get-reports-by-slug)
		1. [[POST] /reports](#post-reports)
		1. [[PUT] /reports](#put-reports)
		1. [[DELETE] /reports](#delete-reports)
		1. [[POST] /reports/batch](#post-reportsbatch)
1. [Settings](#settings)
	1. [dashboard](#dashboard)
		1. [[GET] /dashboards](#get-dashboards)
		1. [[GET] /dashboards/pages](#get-dashboardspages)
		1. [[GET] /dashboards by ID](#get-dashboards-by-id)
		1. [[GET] /dashboards by slug](#get-dashboards-by-slug)
		1. [[POST] /dashboards](#post-dashboards)
		1. [[PUT] /dashboards](#put-dashboards)
		1. [[DELETE] /dashboards](#delete-dashboards)
		1. [[POST] /dashboards/batch](#post-dashboardsbatch)
	1. [query](#query)
		1. [[GET] /queries](#get-queries)
		1. [[GET] /queries/pages](#get-queriespages)
		1. [[GET] /queries by ID](#get-queries-by-id)
		1. [[POST] /queries](#post-queries)
		1. [[PUT] /queries](#put-queries)
		1. [[DELETE] /queries](#delete-queries)
		1. [[POST] /queries/batch](#post-queriesbatch)
	1. [report](#report)
		1. [[GET] /reports](#get-reports)
		1. [[GET] /reports/pages](#get-reportspages)
		1. [[GET] /reports by ID](#get-reports-by-id)
		1. [[GET] /reports by slug](#get-reports-by-slug)
		1. [[POST] /reports](#post-reports)
		1. [[PUT] /reports](#put-reports)
		1. [[DELETE] /reports](#delete-reports)
		1. [[POST] /reports/batch](#post-reportsbatch)
1. [Variable testing](#variable-testing)
1. [[GET] Health check](#get-health-check)

## 0.auth
[Back to Contents](#contents)

### [POST] /login
[Back to 0.auth](#0auth)

**HTTP method**: POST
**Authentication type**: noauth
**Url**: `none to be found.`

Description: Starts a session in the system, necessary for any sort of operations, since it generates the auth token that you'll be using

This request does not have any query params available, or documented.

This request does not have an example request body.

### [POST] /logout
[Back to 0.auth](#0auth)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `none to be found.`

Description: Finishes a session, you may want to restart, or you won't be using that session any longer, in which case, it'd be a good practice to stop the session.

This request does not have any query params available, or documented.

This request does not have an example request body.

## Analytics
[Back to Contents](#contents)

All of the requests related to the analytics services endpoint. Mainly focused on providing data for queries and "business" reports.

`/analytics`

### Dashboard
[Back to Analytics](#analytics)

A set of queries, reports and visualizations so that multiple data can be centralized into one source of information for the final user

`/dashboards`

#### [GET] /dashboards
[Back to Dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false&search&sort=ASC`

Description: Gets all of the elements from the entity, allowing for pagination (strongly recomended).

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |
| **search** |  | The string to filter by |
| **sort** | ASC | How will it be sorted, Ascendant by default. ASC - Ascendant DESC - Descendant |

This request does not have an example request body.

#### [GET] /dashboards/pages
[Back to Dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route/pages`

Description: Gets all the amount of pages for the entity

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by ID
[Back to Dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route/1`

Description: Returns an element by it's ID.

| **field** | **value** | **required** |
| --- | --- | --- |
| ***[HTTP]*** **ID** | The element's ID | **yes** |

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by slug
[Back to Dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route/cool-dashboard`

Description: Returns an element by it's slug.

| **field** | **value** | **required** |
| --- | --- | --- |
| ***[HTTP]*** **slug** | The element's slug | **yes** |

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards
[Back to Dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `/route`

Description: Creates a new element

| **field** | **value** | **required** |
| --- | --- | --- |
| **name** | The name, or title, of the dashboard | **yes** |
| **description** | The description of it's purpose of whatever the user may deem important | **no** |
| **slug** | The verbose identification of the dashboard, must be unique | **no, will be generated by default** |
| **visualizations** | An array of visualizations that the dashboard may use | **no** |

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "name": "Best students performance from 2020-2021",
    "description": "An analysis of the students' performance while the COVID-19 was was spreaded to evaluate the usefulness of the platform",
    "slug": "2020-21-best-students",
    "visualizations": [
        {
            "id": 1,
            "title": "Special KPIs",
            "description": "Metrics to analyze the performance",
            "width": 340,
            "height": 640,
            "x": 0,
            "y": 0
        },
        {
            "id": 2,
            "title": "Special KPIs",
            "description": "Metrics to analyze the performance",
            "width": 340,
            "height": 640,
            "x": 0,
            "y": 0
        },
        {
            "id": 3,
            "title": "Special KPIs",
            "description": "Metrics to analyze the performance",
            "width": 340,
            "height": 640,
            "x": 0,
            "y": 0
        },
        {
            "id": 4,
            "title": "Special KPIs",
            "description": "Metrics to analyze the performance",
            "width": 340,
            "height": 640,
            "x": 0,
            "y": 0
        }
    ]
}
```
</details>

#### [PUT] /dashboards
[Back to Dashboard](#dashboard)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `/route/1`

Description: Modifies an element.

| **field** | **value** | **required** |
| --- | --- | --- |
| ***[HTTP]*** **ID** | The element's ID | **yes** |
| **fieldToUpdate** | The new value | **At least one should be sent** |

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "name": "My new dashboard name"
}
```
</details>

#### [DELETE] /dashboards
[Back to Dashboard](#dashboard)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `/route/1`

Description: Deletes an element.

| **field** | **value** | **required** |
| --- | --- | --- |
| ***[HTTP]*** **ID** | The element's ID | **yes** |

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards/batch
[Back to Dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `/route/batch`

Description: Creates, updates and/or deletes a bunch of elements in batch.

| **field** | **value** | **required** |
| --- | --- | --- |
| **create** | All of the elements to create, without ID | **No** |
| **update** | All of the elements to update, with ID | **No** |
| **delete** | All of the elements to delete, just the ID | **No** |

While none of them are required by default, at least one of them should be passed for the endpoint to compute, otherwise, an error may be thrown.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "create": [
        {}
    ],
    "update": [
        {}
    ],
    "delete": [
        1,
        2,
        3
    ]
}
```
</details>

### Query
[Back to Analytics](#analytics)

#### [GET] /queries
[Back to Query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /queries/pages
[Back to Query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /queries by ID
[Back to Query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries
[Back to Query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /queries
[Back to Query](#query)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /queries
[Back to Query](#query)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries/batch
[Back to Query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### Report
[Back to Analytics](#analytics)

#### [GET] /reports
[Back to Report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /reports/pages
[Back to Report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by ID
[Back to Report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by slug
[Back to Report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports
[Back to Report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /reports
[Back to Report](#report)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /reports
[Back to Report](#report)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports/batch
[Back to Report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

## Courses
[Back to Contents](#contents)

### dashboard
[Back to Courses](#courses)

#### [GET] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /dashboards/pages
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by ID
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by slug
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards/batch
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### query
[Back to Courses](#courses)

#### [GET] /queries
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /queries/pages
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /queries by ID
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /queries
[Back to query](#query)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /queries
[Back to query](#query)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries/batch
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### report
[Back to Courses](#courses)

#### [GET] /reports
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /reports/pages
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by ID
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by slug
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /reports
[Back to report](#report)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /reports
[Back to report](#report)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports/batch
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

## Quizzes
[Back to Contents](#contents)

### dashboard
[Back to Quizzes](#quizzes)

#### [GET] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /dashboards/pages
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by ID
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by slug
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards/batch
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### query
[Back to Quizzes](#quizzes)

#### [GET] /queries
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /queries/pages
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /queries by ID
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /queries
[Back to query](#query)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /queries
[Back to query](#query)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries/batch
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### report
[Back to Quizzes](#quizzes)

#### [GET] /reports
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /reports/pages
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by ID
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by slug
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /reports
[Back to report](#report)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /reports
[Back to report](#report)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports/batch
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

## Roles
[Back to Contents](#contents)

### dashboard
[Back to Roles](#roles)

#### [GET] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /dashboards/pages
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by ID
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by slug
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards/batch
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### query
[Back to Roles](#roles)

#### [GET] /queries
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /queries/pages
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /queries by ID
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /queries
[Back to query](#query)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /queries
[Back to query](#query)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries/batch
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### report
[Back to Roles](#roles)

#### [GET] /reports
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /reports/pages
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by ID
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by slug
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /reports
[Back to report](#report)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /reports
[Back to report](#report)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports/batch
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

## Settings
[Back to Contents](#contents)

### dashboard
[Back to Settings](#settings)

#### [GET] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /dashboards/pages
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by ID
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /dashboards by slug
[Back to dashboard](#dashboard)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /dashboards
[Back to dashboard](#dashboard)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /dashboards/batch
[Back to dashboard](#dashboard)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### query
[Back to Settings](#settings)

#### [GET] /queries
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /queries/pages
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /queries by ID
[Back to query](#query)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /queries
[Back to query](#query)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /queries
[Back to query](#query)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /queries/batch
[Back to query](#query)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

### report
[Back to Settings](#settings)

#### [GET] /reports
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `/route?page=1&perPage=10&indexStartsAt=0&noPagination=false`

Description: None.

**HTTP query params**

| **key** | **value** | **description** |
| --- | --- | --- |
| **page** | 1 | The page to request, if none given, it will be the first by default, if it exceeds the maximum, it will return the last page instead of an error |
| **perPage** | 10 | The amount of results per page, 10 by default |
| **indexStartsAt** | 0 | The starting page index, it will be 1 by default |
| **noPagination** | false | Return ALL elements found, heavy load and risky, should require higher access levels, it will be false by default |

This request does not have an example request body.

#### [GET] /reports/pages
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/pages`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by ID
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [GET] /reports by slug
[Back to report](#report)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/1`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{}
```
</details>

#### [PUT] /reports
[Back to report](#report)

**HTTP method**: PUT
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [DELETE] /reports
[Back to report](#report)

**HTTP method**: DELETE
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

#### [POST] /reports/batch
[Back to report](#report)

**HTTP method**: POST
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080/route/batch`

Description: None.

This request does not have any query params available, or documented.

<details>
<summary>Show request's body</summary>
  
```json
{
    "elements": [
        {}
    ]
}
```
</details>

## Variable testing
[Back to Contents](#contents)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.

## [GET] Health check
[Back to Contents](#contents)

**HTTP method**: GET
**Authentication type**: none to be kown.
**Url**: `http://localhost:8080`

Description: None.

This request does not have any query params available, or documented.

This request does not have an example request body.