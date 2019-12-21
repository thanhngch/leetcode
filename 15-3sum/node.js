/**
 * @param {number[]} nums
 * @return {object}
 */
var indexArr = function(nums) {
    let map = {};
    for (let i = 0; i < nums.length; i++) {
        if (map[nums[i]] === undefined) {
            map[nums[i]] = [i];
        } else {
            map[nums[i]].push(i);
        }
    }
    return map;
};

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    // sort nums
    nums = nums.sort((a, b) => (a - b));
    if (nums.length < 3) {
        return [];
    }
    let triple = [];
    for (let i = 2; i < nums.length; i++) {
        if (nums[i] === 0 && nums[i-1] === 0 && nums[i-2] === 0) {
            triple.push([0,0,0]);
            break;
        }
    }
    for (let i = 2; i < nums.length; i++) {
        if (nums[i] === nums[i-1] && nums[i] === nums[i-2]) {
            nums.splice(i, 1);
            i--;
        }
    }
    let map = indexArr(nums);
    for (let i = 0; i < nums.length - 2; i++) {
        if (nums[i] > 0) {
            break;
        }
        for (let j = i + 1; j < nums.length - 1; j++) {
            let otherNumber = -(nums[i] + nums[j]);
            if (map[otherNumber] !== undefined) {
                for (let k = 0; k < map[otherNumber].length; k++) {
                    if (map[otherNumber][k] > j) {
                        // check triple is exists
                        let numsSorted = [nums[i], nums[j], otherNumber];
                        let isExist = false;

                        if (triple.length > 0) {
                            if (triple[triple.length-1][0] === numsSorted[0]) {
                                for (let l = triple.length - 1; l >= 0; l--) {
                                    if (triple[l][0] !== numsSorted[0]) {
                                        break;
                                    } else {
                                        if (triple[l][1] === numsSorted[1]
                                            && triple[l][2] === numsSorted[2]) {
                                                isExist = true;
                                        }
                                    }
                                }
                            }
                        }
                        if (!isExist) {
                           // add to tripple
                            triple.push([nums[i], nums[j], otherNumber]);
                        }
                        break;
                    }
                }
            }
        }
    }
    return triple;
};

var main = function() {
    // let nums = [-1, 0, 1, 2, -1, -4];
    // let nums = [-1, 0, 1, 2, -1, -4];
    // let nums = [-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6];
    let nums = [0,-5,3,-4,1,3,-4,-2,-2,-2,0,3,0,1,-4,-2,0];
    // nums = nums.sort((a, b) => (a - b));
    // console.log(nums);
    // for (let i = 2; i < nums.length; i++) {
    //     if (nums[i] === nums[i-1] && nums[i] === nums[i-2]) {
    //         nums.splice(i, 1);
    //         i--;
    //     }
    // }
    // console.log(nums);
    // console.log(indexArr(nums));
    console.log(threeSum(nums));
}

main();
