with open('../input.txt', 'r') as file:
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

safe_reports = 0
for report in reports: 
    is_safe = True 
    if not (is_increasing(report) or is_decreasing(report)):
        is_safe = False
    else:
        for i in range(len(report) - 1):
            if not differs(report[i], report[i + 1]):
                is_safe = False
                break

    if is_safe: 
        safe_reports += 1
            
print(safe_reports) 