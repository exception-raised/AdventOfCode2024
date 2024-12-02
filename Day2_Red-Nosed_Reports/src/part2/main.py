with open('../../input.txt', 'r') as file:
    lines = file.readlines()

reports = []

for line in lines:
    array = list(map(int, line.strip().split()))
    reports.append(array)

def is_increasing(arr):
    for i in range(len(arr) - 1):
        if arr[i] >= arr[i + 1]:
            return False
    return True

def is_decreasing(arr):
    for i in range(len(arr) - 1):
        if arr[i] <= arr[i + 1]:
            return False
    return True

def differs(x, y):
    diff = abs(x - y)
    return 1 <= diff <= 3

def is_safe(report):
    if is_increasing(report) or is_decreasing(report):
        for i in range(len(report) - 1):
            if not differs(report[i], report[i + 1]):
                return False
        return True
    return False

safe_reports = 0
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
