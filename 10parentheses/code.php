<?php

class Stack {
    var $data = [];
    function push($value) {
        array_push($this->data, $value);
    }
    function pop() {
        array_pop($this->data);
    }
    function top() {
        if (count($this->data) == 0) {
            return '0';
        }
        return $this->data[count($this->data) - 1];
    }
    function isEmpty() {
        return count($this->data) == 0;
    }
}
class Solution {
    function validChar($c) {
        switch ($c) {
        case '{':
        case '}':
        case '[':
        case ']':
        case '(':
        case ')':
            return true;
        }
        return false;
    }
    function invertChar($c) {
        switch ($c) {
        case '{':
            return '}';
        case '(':
            return ')';
        case '[':
            return ']';
        }
        return '0';
    }
    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
        $stack = new Stack();
        for ($i = 0; $i < strlen($s); $i++) {
            $c = $s[$i];
            if ($this->validChar($c)) {
                if ($this->invertChar($stack->top()) == $c) {
                    $stack->pop();
                } else {
                    $stack->push($c);
                }
            } else {
                return false;
            }
        }
        return $stack->isEmpty();
    }
}

// printf("\n Is Valid: %s", (new Solution())->isValid('()') == true ? 'true' : 'false');
// Test
class Cases {
    var $in;
    var $out;
}

function setupData() {
    $fileContent = file_get_contents('./test.txt', true);

    $dataArray = explode("\n", $fileContent);
    $cases = [];
    foreach ($dataArray as $line) {
        $lineSplit = explode(" ", $line);
        $case = new Cases();
        $case->in = $lineSplit[0];
        $case->out = $lineSplit[1];
        $cases[] = $case;
    }
    return $cases;
}

function test() {
    $cases = setupData();
    $solution = new Solution();
    $t1 = microtime(true);
    foreach($cases as $case) {
        $got = $solution->isValid($case->in);
        if ($got != $case->out) {
            printf("Test %s got %s, want %s\n", $case->in, $got, $case->out);
        }
    }
    $t2 = microtime(true);
    $elaps = ($t2 - $t1) * 1E6;
    printf("Finish test in %.2f μs\n", $elaps);
    printf("Memory is using %.2fMB\n", memory_get_usage()/1024/1024);
    printf("Memory allocated from system %.2fMB\n", memory_get_usage(true)/1024/1024);
}

test();
