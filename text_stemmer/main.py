from aiohttp import web

from stemmer import Stemmer

_stemmer = Stemmer()


async def handle(request):
    text = await _stemmer.apply(await request.text())
    return web.Response(text=text)


app = web.Application()
app.add_routes([web.post('/stemmer', handle)])

if __name__ == '__main__':
    web.run_app(app, port=8081)
