<?php

class Solution {

    /**
     * @param String $haystack
     * @param String $needle
     * @return Integer
     */
    function strStr($haystack, $needle) {
        if ($needle === '') {
            return 0;
        }
        $pos = strpos($haystack, $needle);
        if ($pos === FALSE) {
            return -1;
        }
        return $pos;
    }
}

$solution = new Solution();
print($solution->strStr('hello', 'll') . "\n");
print($solution->strStr('123','3') . "\n");
print($solution->strStr('123','1') . "\n");
print($solution->strStr('','') . "\n");
print($solution->strStr('aaaa', 'z') . "\n");
