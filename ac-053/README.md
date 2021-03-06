# 剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。

## 解题思路：
1.暴力循环
2.二分法

# 剑指 Offer 53 - II. 0～n-1中缺失的数字
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:

输入: [0,1,3]
输出: 2

示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8

## 解题思路

### 思路一：二分查找
定义 left:=0 right:=len(nums) 用来确定数组的 mid
因为nums是有序数组，如果mid下标的值和mid不相同就在左边查找
如果 nums[mid]==mid ，说明左边是连续的有序数组，缺失的数字就在右边查找
left 需要向上取整+1

### 思路二：总和差值
- 1.计算当数组不缺数字时的总和 count
- 2.count减去该数组nums的总和
- 3.得到的差值即是缺失的数字

