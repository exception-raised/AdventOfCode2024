import unittest
import sys
import os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))
from shared import differs, is_increasing, is_decreasing, parse_input_into_arrays, is_safe

class TestRedNosedReports(unittest.TestCase):
    def test_differs(self):
        self.assertTrue(differs(5, 7))  # Difference is 2
        self.assertTrue(differs(7, 5))  # Difference is 2
        self.assertFalse(differs(5, 9))  # Difference is 4
        self.assertFalse(differs(5, 5))  # Difference is 0

    def test_is_increasing(self):
        self.assertTrue(is_increasing([1, 2, 3, 4, 5]))
        self.assertFalse(is_increasing([5, 4, 3, 2, 1]))
        self.assertFalse(is_increasing([1, 3, 2, 4, 5]))
        self.assertFalse(is_increasing([1, 1, 2, 3, 4])) 

    def test_is_decreasing(self):
        self.assertTrue(is_decreasing([5, 4, 3, 2, 1]))
        self.assertFalse(is_decreasing([1, 2, 3, 4, 5]))
        self.assertFalse(is_decreasing([5, 3, 4, 2, 1]))
        self.assertFalse(is_decreasing([5, 5, 4, 3, 2])) 

    def test_parse_input_into_arrays(self):
        with open("mock_input.txt", "w") as file:
            file.write("1 2 3 4 5\n")
            file.write("5 4 3 2 1\n")
            file.write("1 3 2 4 5\n")

        reports = parse_input_into_arrays("mock_input.txt")
        self.assertEqual(reports, [
            [1, 2, 3, 4, 5],
            [5, 4, 3, 2, 1],
            [1, 3, 2, 4, 5]
        ])

    def test_is_safe(self):
        self.assertTrue(is_safe([1, 2, 3, 4, 5]))
        self.assertTrue(is_safe([5, 4, 3, 2, 1]))  
        self.assertFalse(is_safe([1, 3, 2, 4, 5])) 
        self.assertFalse(is_safe([1, 1, 2, 3, 4])) 
        self.assertFalse(is_safe([5, 9, 4, 3, 2]))  

if __name__ == "__main__":
    unittest.main()
