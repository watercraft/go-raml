from sanic import Blueprint
from sanic.views import HTTPMethodView
from sanic.response import text
import deliveries_api


deliveries_if = Blueprint('deliveries_if')


class deliveriesView(HTTPMethodView):

    async def get(self, request):

        return await deliveries_api.deliveries_get(request)

    async def post(self, request):

        return await deliveries_api.deliveries_post(request)


deliveries_if.add_route(deliveriesView.as_view(), '/deliveries')


class deliveries_bydeliveryIdView(HTTPMethodView):

    async def get(self, request, deliveryId):

        return await deliveries_api.deliveries_byDeliveryId_get(request, deliveryId)

    async def patch(self, request, deliveryId):

        return await deliveries_api.deliveries_byDeliveryId_patch(request, deliveryId)

    async def delete(self, request, deliveryId):

        return await deliveries_api.deliveries_byDeliveryId_delete(request, deliveryId)


deliveries_if.add_route(deliveries_bydeliveryIdView.as_view(), '/deliveries/<deliveryId>')
