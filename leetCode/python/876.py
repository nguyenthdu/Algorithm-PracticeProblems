class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def middle_node(head):
    slow = fast = head

    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

    return slow

def print_from_middle_to_end(head):
    middle = middle_node(head)
    current = middle

    while current:
        print(current.val, end=" ")
        current = current.next
    print()

if __name__ == "__main__":
    # Tạo một danh sách ví dụ
    node1 = ListNode(1)
    node2 = ListNode(2)
    node3 = ListNode(3)
    node4 = ListNode(4)
    node5 = ListNode(5)

    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node5

    # In ra danh sách từ giữa đến hết
    print("List from middle to end:")
    print_from_middle_to_end(node1)
