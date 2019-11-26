<?php

class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function reverse($x) {
        $n = 0;
        while ($x != 0) {
            $n *= 10;
            $n += $x % 10;
            $x = ($x - ($x%10))/10;
        }
        if ($n > 2147483647 || $n < -2147483648) {
            return 0;
        }
        return $n;
    }
}

echo (new Solution())->reverse(-123);
