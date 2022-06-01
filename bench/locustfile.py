from datetime import datetime, timedelta
import string
from locust import HttpUser, task
import random

PET_KINDS = ["Cat", "Dog", "Hamster", "Guinea Pig", "Alligator"]
PET_BREEDS = ["British Shorthair", "Labrador Retriever"]
ACCESS_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQwNzAwOTV9.ul5Cx40cergadeFrR_3_Ug7ljgLNZciIc6NjbVHlQCU"

def random_string():
    return "".join(random.choice(string.ascii_lowercase) for i in range(10))

def random_date():
    start = datetime.strptime("2010-01-01", "%Y-%m-%d")
    end = datetime.strptime("2022-01-01", "%Y-%m-%d")
    delta = end - start 
    int_delta = (delta.days * 24 * 60 * 60) + delta.seconds
    random_second = random.randrange(int_delta)
    return start + timedelta(seconds=random_second)

class TestUser(HttpUser):
    @task
    def add_users(self):
        body = {
            "name": random_string(),
            "phoneNumber": "+380" + str(random.randint(100000000, 999999999))
        }
        with self.client.post("/users", json=body, catch_response=True) as response:
            if response.status_code != 201:
                response.failure("Failed to create user: " + response.text)
    
    @task
    def add_pet(self):
        body = {
            "name": random_string(),
            "ownerId": random.randint(1, 1),
            "kind": random.choice(PET_KINDS),
            "breed": random.choice(PET_BREEDS),
            "birthDate": random_date().strftime("%Y-%m-%d")
        }
        headers = {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + ACCESS_TOKEN
        }
        with self.client.post("/pets", json=body, headers=headers, catch_response=True) as response:
            if response.status_code != 201:
                response.failure("Failed to create pet: " + str(body))
    
    @task
    def get_users(self):
        self.client.get("/users")

    @task
    def get_pets(self):
        self.client.get("/pets")