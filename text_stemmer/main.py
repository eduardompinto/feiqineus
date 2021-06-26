from fastapi import FastAPI, Request

from stemmer import Stemmer

_stemmer = Stemmer()

app = FastAPI()


@app.post("/stemmer")
async def handle(req: Request):
    text_to_parse = await req.body()
    text = await _stemmer.apply(text_to_parse.decode("utf-8"))
    return text

