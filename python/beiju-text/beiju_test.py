"""
/* UVA11988 Broken Keyboard (a.k.a. Beiju Text)  */
"""
import unittest
from beiju import parse


class BeijuTest(unittest.TestCase):
    test_cases = [
        {
            "name": "default",
            "sentences": "This_is_a_[Beiju]_text",
            "expected": "BeijuThis_is_a__text"
        },
        {
            "name": "test case 1",
            "sentences": "[[]][][]Happy_Birthday_to_Tsinghua_University",
            "expected": "Happy_Birthday_to_Tsinghua_University"
        },
        {
            "name": "test case 2",
            "sentences": "abc[def]ed[ed]]d[ddee",
            "expected": "ddeeeddefabcedd"
        }
    ]

    def test_beiju(self):
        for case in self.test_cases:
            self.assertEqual(parse(case['sentences']), case['expected'])
