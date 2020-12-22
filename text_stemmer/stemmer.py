import os
import unicodedata

from nltk.corpus import stopwords
from nltk.data import path
from nltk.stem import SnowballStemmer


class Stemmer:

    def __init__(self):
        path.append(os.path.abspath('./nltk_data'))
        self._stemmer = SnowballStemmer("portuguese")
        self._stopwords = stopwords.words('portuguese')

    async def apply(self, sentence: str):
        return await self._portuguese_stemming(sentence)

    async def _portuguese_stemming(self, sentence):
        sentence = await self._clean_sentence(sentence)
        stem = ' '.join(self._stemmer.stem(word) for word in sentence.split(' '))
        stem = sorted(stem.split(' '))
        return ' '.join(stem)

    async def _clean_sentence(self, sentence):
        clean_sentence = ''.join(ch for ch in sentence if ch.isalnum() or ch == ' ')
        normalized_sentence = str(
            unicodedata.normalize('NFKD', clean_sentence).encode('ASCII', 'ignore'),
            'UTF-8'
        )

        def valid_word(w: str) -> bool:
            return w.lower() not in self._stopwords and w != ' '

        return ' '.join(
            [w for w in normalized_sentence.split() if valid_word(w)]
        )
