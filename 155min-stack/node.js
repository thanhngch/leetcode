/**
 * initialize your data structure here.
 */
var MinStack = function() {
    this.val = [];
    this.min = null;
};

/** 
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
    this.val.push(x);
    if (this.min !== null) {
        this.min = Math.min(this.min, x);
    } else {
        this.min = x;
    }
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    if (this.top() === this.min && this.min !== null) {
        this.val.pop();
        if (this.val.length > 0) {
            let min = this.val.reduce(function (a, b) {
                return ( a < b ? a : b );
            });
            this.min = min;
        } else {
            this.min = null;
        }
    } else {
        this.val.pop();
    }
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    if (this.val.length > 0) {
        return this.val[this.val.length - 1];
    }
    return null;
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
    return this.min;
};

/** 
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */

var minStack = new MinStack();
minStack.push(0);
minStack.push(1);
minStack.push(2);
minStack.push(3);
console.log(minStack);
console.log(minStack.getMin()); //  --> Returns 0.
minStack.pop();
minStack.pop();
minStack.pop();
console.log(minStack.top());    //  --> Returns null.
console.log(minStack.getMin()); //  --> Returns null.
