# NusaMeals

adalah website penyedia makanan tradisional di indonesia

API product

Create
Req
{
    "name": 
    "stock":
    "type":
}
Res
{
    "id"
    "name":
    "stock":
    "type":
}

Update :id
{
    "id"
    "name":
    "stock":
    "type":
}


DeleteByID :id -> req
res
{
    "message"
}

GetAll 
[
    {
        "id"
        "name":
        "stock":
        "type":
    },
    {
        "id"
        "name":
        "stock":
        "type":
    }
]

GetById :id
{
    "id"
    "name":
    "stock":
    "type":
}

GetAllByType :type
[
    {
        "id"
        "name":
        "stock":
        "type":
    },
    {
        "id"
        "name":
        "stock":
        "type":
    }
]