import unittest


class BeijuTest(unittest):
    test_cases = [
        {
            "name": "",
            "input": "This_is_a_[Beiju]_text",
            "expected": "BeijuThis_is_a__text"
        },
        {
            "name": "",
            "input": "[[]][][]Happy_Birthday_to_Tsinghua_University",
            "expected": "Happy_Birthday_to_Tsinghua_University"
        }
    ]

    def test_beiju(self):
        for test_case in self.test_cases:
            # assert test_case["expected"] == test_case["input"]
            pass
