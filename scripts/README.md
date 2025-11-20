"""
API Service Project
=====================

This project provides a simple REST API using Flask.

Installation
------------

To install the required packages, run:

```bash
pip install -r requirements.txt
```

Running the Application
------------------------

To run the application, execute the following command:

```bash
flask run
```

API Endpoints
-------------

The API provides the following endpoints:

*   **GET /users**: Retrieves a list of all users.
*   **GET /users/:id**: Retrieves a user by ID.
*   **POST /users**: Creates a new user.
*   **PUT /users/:id**: Updates an existing user.
*   **DELETE /users/:id**: Deletes a user by ID.

User Model
------------

The user model is defined as follows:

```python
class User:
    def __init__(self, id, name, email):
        self.id = id
        self.name = name
        self.email = email

    def __repr__(self):
        return f"User('{self.id}', '{self.name}', '{self.email}')"
```

API Controllers
----------------

The API controllers are defined in the `controllers` module. Each controller handles a specific endpoint and handles requests and responses accordingly.

```python
from flask import jsonify, request
from models import User

class UserController:
    def get_users(self):
        users = User.query.all()
        return jsonify([user.to_dict() for user in users])

    def get_user(self, id):
        user = User.query.get(id)
        if user:
            return jsonify(user.to_dict())
        return jsonify({"error": "User not found"}), 404

    def create_user(self):
        data = request.get_json()
        user = User(**data)
        db.session.add(user)
        db.session.commit()
        return jsonify(user.to_dict()), 201

    def update_user(self, id):
        user = User.query.get(id)
        if user:
            data = request.get_json()
            for key, value in data.items():
                setattr(user, key, value)
            db.session.commit()
            return jsonify(user.to_dict())
        return jsonify({"error": "User not found"}), 404

    def delete_user(self, id):
        user = User.query.get(id)
        if user:
            db.session.delete(user)
            db.session.commit()
            return jsonify({"message": "User deleted successfully"})
        return jsonify({"error": "User not found"}), 404
```

Models
-------

The user model is defined in the `models` module.

```python
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy(app)

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    email = db.Column(db.String(100), nullable=False, unique=True)

    def to_dict(self):
        return {"id": self.id, "name": self.name, "email": self.email}
```

Views
-----

The API views are defined in the `views` module. Each view handles a specific endpoint and handles requests and responses accordingly.

```python
from flask import Blueprint
from controllers import user_controller

user_blueprint = Blueprint('users', __name__)

user_blueprint.add_url_rule('/users', view_func=user_controller.get_users, methods=['GET'])
user_blueprint.add_url_rule('/users/<int:id>', view_func=user_controller.get_user, methods=['GET'])
user_blueprint.add_url_rule('/users', view_func=user_controller.create_user, methods=['POST'])
user_blueprint.add_url_rule('/users/<int:id>', view_func=user_controller.update_user, methods=['PUT'])
user_blueprint.add_url_rule('/users/<int:id>', view_func=user_controller.delete_user, methods=['DELETE'])