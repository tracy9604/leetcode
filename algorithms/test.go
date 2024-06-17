package algorithms
def insert(intervals, newInterval):
result = []
i = 0

# Add all intervals that end before the new interval starts
while i < len(intervals) and intervals[i][1] < newInterval[0]:
result.append(intervals[i])
i += 1

# Merge all overlapping intervals to one considering the new interval
while i < len(intervals) and intervals[i][0] <= newInterval[1]:
newInterval[0] = min(newInterval[0], intervals[i][0])
newInterval[1] = max(newInterval[1], intervals[i][1])
i += 1

# Add the merged interval
result.append(newInterval)

# Add all the rest
while i < len(intervals):
result.append(intervals[i])
i += 1

return result

