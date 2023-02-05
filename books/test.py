import requests

# url = "http://localhost:8080/books"
url = "https://gobooks.abooishaaq.repl.co/books"

def add_book() -> str:
    book = {
        "title": "The Alchemist",
        "author": "Paulo Coelho",
    }
    response = requests.post(url, json=book)
    res = response.json()
    print(res)
    return res["ID"]


def get_books():
    response = requests.get(url)
    print(response.json())

def update_book(id: str):
    book = {
        "title": "The Alchemist",
        "author": "Ammar",
    }
    response = requests.post(f"{url}/{id}", json=book)
    print(response.json())

def delete_book(id: str):
    response = requests.delete(f"{url}/{id}")
    print(response.json())

if __name__ == "__main__":
    id = add_book()
    get_books()
    update_book(id)
    get_books()
    delete_book(id)
    get_books()

