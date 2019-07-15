from app import app
from app import api 
from flask import request
from flask_restful import Resource


class HelloWorld(Resource):
    def get(self):
        return 'Hello World'

api.add_resource(HelloWorld, '/channel/list')
