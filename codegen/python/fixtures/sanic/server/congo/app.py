from sanic import Sanic
from sanic.response import json

from deliveries_if import deliveries_if
from drones_if import drones_if

import os
dir_path = os.path.dirname(os.path.realpath(__file__))

app = Sanic(__name__)

app.blueprint(deliveries_if)
app.blueprint(drones_if)


app.static('/apidocs', dir_path + '/apidocs')
app.static('/', dir_path + '/index.html')

if __name__ == "__main__":
    app.run(debug=True, port=5000, workers=2)
