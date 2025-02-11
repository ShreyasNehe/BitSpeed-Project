# Identity Reconciliation

## Bitespeed Backend Task: Identity Reconciliation

### **Tech Stack Used**
- **Programming Language**: Golang
- **Frameworks & Packages**: Gorilla Mux, GORM
- **Database**: PostgreSQL
- **Deployment**: Render.com

### **Base URL**
```
https://identity-reconciliation-gq7z.onrender.com/
```

### **Service Route**
```
/bitespeed/identity_reconciliation/v1
```

---

## **API Endpoints**

### **1. Health Check API**
Endpoint to check if the service is running.

---

### **2. Get All Contacts**
#### **Endpoint:**
```
GET /contact/getAll
```

#### **Reference cURL:**
```sh
curl --location 'https://identity-reconciliation-gq7z.onrender.com/bitespeed/identity_reconciliation/v1/contact/getAll' \
--header 'Content-Type: application/json' \
--data '{
  "start": 0,
  "end": -1
}'
```

#### **Response:**
```json
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
```

---

### **3. Identify Contacts**
#### **Endpoint:**
```
POST /contact/identify
```

#### **Reference cURL:**
```sh
curl --location 'https://identity-reconciliation-gq7z.onrender.com/bitespeed/identity_reconciliation/v1/contact/identify' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "george@hillvalley.edu",
  "phoneNumber": "717171"
}'
```

#### **Response:**
```json
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
```

---

## **Running the Project Locally**

### **Step 1: Clone the Repository**
```sh
git clone https://github.com/ShreyasNehe/BitSpeed-Project.git
```

### **Step 2: Configure Environment Variables**
Navigate to the project directory:
```sh
cd identity_reconciliation
```
Create and edit the `Config.env` file:
```sh
vi Config.env
```
Add the following configuration (replace with actual values):
```
PSQL_DB_URL=<your_database_url>
PORT=4993
```

### **Step 3: Start the Server**
```sh
go run main.go
```

---

## **Contact Information**
- **Name**: Shreyas Nehe
- **LinkedIn**:(https://www.linkedin.com/in/shreyasnehe/)
- 
- **Email**: shreyas.nehe2021@gmail.com
- **Skills**: Golang, Python, MongoDB, MySQL, Kafka, Redis, JavaScript, ReactJS,NextJs,ExprresJS,AWS,Oracle
- **Interest**: Passionate about learning new technologies

