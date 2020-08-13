# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def insertionSortList(self, head: ListNode) -> ListNode:
        if not head:                            
            return head
        dummy=ListNode(float(-inf))             #首项负无穷保证插的数据不在最前头
        dummy.next=head                         #用于返回答案
        pre=head                                    
        nxt=head.next
        while nxt:                              #最后一项为None结束循环
            if nxt.val>pre.val:                 #从头开始，如果后一项大于前一项不改变节点
               pre=nxt                          #这两行移动指针（往后挪一格）
               nxt=nxt.next
            else:                               #如果后项小需要把后项nxt的节点插到正确位置
                pre.next=nxt.next               #把nxt指向的节点拿出来
                cp1=dummy                       #这两个指针负责从头开始比较nxt的位置
                cp2=dummy.next                  #使用dummy的原因见11行注释
                while nxt.val>cp2.val:          #因为前面是排好序的循环结束nxt正好在cp1和cp2中间
                    cp1=cp2                     #
                    cp2=cp2.next                #
                nxt.next=cp2                    #这两行负责插
                cp1.next=nxt                    #把nxt指向的节点查到cp1和cp2中间
                nxt=pre.next                    #指针从哪来回哪去 准备下一个循环
        return dummy.next

