# identity_reconciliation
Bitespeed Backend Task: Identity Reconciliation

Tech Stack Used 
Programming - Golang
packages and frameworks - Gorilla Mux, gorm
Databases - PostGreSQL
Deployment - rendor.com
BaseUrl - https://identity-reconciliation-gq7z.onrender.com/
Service Route - bitespeed/identity_reconciliation/v1


Apis--->
Health Check API -->


Get all contacts --->
end point - /contact/getAll
reference curl:
curl --location 'https://identity-reconciliation-gq7z.onrender.com/bitespeed/identity_reconciliation/v1/contact/getAll' \
--header 'Content-Type: application/json' \
--data '{
"start":0,
"end":-1
}'

response:
[
{
"id": 37,
"phoneNumber": "",
"email": "chellas@g.c",
"linkedId": 0,
"linkPrecedence": "primary",
"createdAt": "2024-02-22T13:24:16.324908Z",
"updatedAt": "2024-02-22T13:24:16.324908Z",
"deletedAt": null
}
]

Identify Contacts -
endpoint - /contact/identify
reference curl:
curl --location 'https://identity-reconciliation-gq7z.onrender.com/bitespeed/identity_reconciliation/v1/contact/identify' \
--header 'Content-Type: application/json' \
--data-raw '{
"email":"george@hillvalley.edu",
"phoneNumber": "717171"
}'

response:
{
"contact": {
"primaryContactId": 38,
"emails": [
"george@hillvalley.edu"
],
"phoneNumbers": [
"717171"
],
"secondaryContactIds": null
}
}

Please reach out to me for config data

Name - suresh
LinkedIn - https://www.linkedin.com/in/sureshchalla024/
Naukari - https://www.naukri.com/mnjuser/profile?id=&altresid

Please contact on schalla024@gmail.com for more info 
