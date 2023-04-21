package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// mock func
func isBadVersion(int) bool {
	return true
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	middle := n / 2

	for left <= right {
		isBad := isBadVersion(middle)

		if isBad && !isBadVersion(middle-1) {
			return middle
		}

		if isBad {
			right = middle - 1
		} else {
			left = middle + 1
		}
		middle = (right + left) / 2
	}
	return middle
}
