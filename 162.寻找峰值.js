/*
 * @lc app=leetcode.cn id=162 lang=javascript
 *
 * [162] 寻找峰值
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var findPeakElement = function (nums) {

    const n = nums.length;
    if (n == 1) return 0;

    nums[-1] = nums[n] = -Infinity

    // A[1..n] 中是否存在 i 使得 A[i-1] < A[i] > A[i+1]
    for (let i = 0; i < n; i++) {
        if (nums[i - 1] < nums[i] && nums[i] > nums[i + 1]) {
            return i
        }
    }
};
// @lc code=end

// 哭了 js 没有连续比较运算 连等也有一堆限制
// const nums = [1, 2, 1, 3, 5, 6, 4]
const nums = [1, 2]
console.log(findPeakElement(nums))