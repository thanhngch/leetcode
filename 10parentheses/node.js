/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
    let stack = [];
    for (let i = 0; i < s.length; i++) {
        let c = s[i];
        if (validChar(c)) {
            if (invertChar(topStack(stack)) == c) {
                stack.pop();
            } else {
                stack.push(c);
            }
        } else {
            return false;
        }
    };
    // return stack is empty
    return stack.length == 0;
};
var topStack = function (s) {
    return s[s.length - 1];
}
var validChar = function (c) {
    switch (c) {
        case '{':
        case '}':
        case '[':
        case ']':
        case '(':
        case ')':
            return true
    }
    return false
}

var invertChar = function (c) {
    switch (c) {
        case '{':
            return '}'
        case '(':
            return ')'
        case '[':
            return ']'
    }
}

var getStack = function () {
    var stack = {
        data: []
    };
    stack.push = function (v) {
        this.data.push(v);
    }
    stack.pop = function () {
        this.data.pop();
    }
    stack.top = function () {
        // if (length(this.data) == 0) {
        //     return null
        // }
        return this.data[this.data.length - 1]
    }
    stack.empty = function () {
        return this.data === 0
    }
    return stack;
}
var testStack = function () {
    let stack = getStack();
    stack.push(1);
    stack.push(2);
    console.log(stack.data);
    console.log(stack.top());
    stack.pop();
    console.log(stack.data);
    console.log(stack.top());
}

var main = function () {
    const cases = [{
            in: "()",
            out: true
        },
        {
            in: "[()]",
            out: true
        },
        {
            in: "()[]{}",
            out: true
        },
        {
            in: "(]",
            out: false
        },
        {
            in: "([)]",
            out: false
        },
        {
            in: "{[]}",
            out: true
        },
    ]
    cases.forEach(_case => {
        got = isValid(_case.in);
        if (got != _case.out) {
            console.error(`in ${_case.in}, got ${got} want ${_case.out}`)
        }
    })
    console.log('Finish test')
}

// main();

var test = function() {
    const fs = require('fs');
    const { performance } = require('perf_hooks');

    const fileData = fs.readFileSync('test.txt', 'utf8').split('\n');
    let cases = [];
    fileData.forEach(line => {
        const lineSplit = line.split(' ');
        cases.push({
            in: lineSplit[0],
            out: lineSplit[1]
        })
    });
    const time1 = performance.now();
    cases.forEach(_case => {
        got = isValid(_case.in);
        if (got != _case.out) {
            console.error(`in ${_case.in}, got ${got} want ${_case.out}`)
        }
    });
    const time2 = performance.now();

    console.log(`Finish test in ${((time2 - time1) * 1000).toFixed(2)}µs`);

    const process = require('process');

    const used = process.memoryUsage();
    for (let key in used) {
        console.log(`${key} ${Math.round(used[key] / 1024 / 1024 * 100) / 100} MB`);
    }
}
test();
// testStack();