
class Solution:
    def checkOverlap(self, radius: int, x_center: int, y_center: int, x1: int, y1: int, x2: int, y2: int) -> bool:
        # 条件 1：首先判断圆心是否在矩形内
        if x1 <= x_center <= x2 and y1 <= y_center <= y2:
            return True
        # 条件 2：圆心位于矩形的上下左右四个区域
        elif x_center > x2 and y1 <= y_center <= y2: # 右
            return radius >= x_center - x2
        elif y_center < y1 and x1 <= x_center <= x2: # 下
            return radius >= y1 - y_center
        elif x_center < x1 and y1<= y_center <= y2: # 左
            return radius >= x1 - x_center
        elif y_center > y2 and x1 <= x_center <= x2: # 上
            return radius >= y_center - y2
        else:
        # 条件 3：判断矩形的四个顶点是否在圆的内部
            return min((x1 - x_center) ** 2 + (y2 - y_center) ** 2,\
                       (x2 - x_center) ** 2 + (y2 - y_center) ** 2, \
                       (x2 - x_center) ** 2 + (y1 - y_center) ** 2, \
                       (x1 - x_center) ** 2 + (y1 - y_center) ** 2) <= radius ** 2


                
radius = 1
x_center = 0
y_center = 0
x1 = 1
y1 = -1
x2 = 3
y2 = 1
res = Solution().checkOverlap(radius, x_center,y_center, x1, y1, x2, y2)
print(res)