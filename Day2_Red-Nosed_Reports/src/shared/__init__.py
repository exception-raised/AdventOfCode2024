def differs(x, y):
    diff = abs(x - y)
    return 1 <= diff <= 3

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

def parse_input_into_arrays(filename: str):
    with open(filename, 'r') as file:
        lines = file.readlines()

    reports = []

    for line in lines:
        array = list(map(int, line.strip().split()))
        reports.append(array)

    return reports

def is_safe(report):
    if is_increasing(report) or is_decreasing(report):
        for i in range(len(report) - 1):
            if not differs(report[i], report[i + 1]):
                return False
        return True
    return False