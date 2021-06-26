import os
import unicodedata

from nltk.corpus import stopwords
from nltk.data import path
from nltk.stem import SnowballStemmer

class Stemmer:

    _PORTUGUESE = 'portuguese'
    _NLTK_DATA = './nltk_data'
    _NFKD = 'NFKD'

    def __init__(self):
        path.append(os.path.abspath(self._NLTK_DATA))
        self._stemmer = SnowballStemmer(self._PORTUGUESE)
        self._stopwords = stopwords.words(self._PORTUGUESE)

    async def apply(self, sentence: str):
        return await self._portuguese_stemming(sentence)

    async def _portuguese_stemming(self, sentence):
        sentence = await self._clean_sentence(sentence)
        stem = sorted(set(self._stemmer.stem(word) for word in sentence.split(' ')))
        return ' '.join(stem)

    async def _clean_sentence(self, sentence):
        normalized_sentence = self._normalize_sentence(sentence=sentence, form=self._NFKD)
        return ' '.join(
            [w for w in normalized_sentence.split() if self._valid_word(w)]
        )

    def _normalize_sentence(self, sentence: str, form: str = _NFKD):
        normalized_sentence = unicodedata.normalize(
            form, self._extract_alnum_and_spaces(sentence)
        ).encode('ASCII', 'ignore')
        return str(normalized_sentence, 'UTF-8')

    def _extract_alnum_and_spaces(self, sentence: str):
        return ''.join(ch for ch in sentence if ch.isalnum() or ch == ' ')

    def _valid_word(self, w: str) -> bool:
        return w.lower() not in self._stopwords and w != ' '
