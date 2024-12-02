import sys
import os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))
from shared import is_safe, parse_input_into_arrays


safe_reports = 0
reports = parse_input_into_arrays('../../input.txt')

for report in reports: 
    if is_safe(report):
        safe_reports += 1
    else:
        for i in range(len(report)):
            modified_report = report[:i] + report[i+1:]
            if is_safe(modified_report):
                safe_reports += 1
                break

print(safe_reports)
