import intersection from 'lodash-es/intersection';
import xor from 'lodash-es/xor';

function isEmpty(aryOrStr) {
  return !aryOrStr.length;
}

function baz(a, b) {
  return a.split(' ').every((text) => b.toLowerCase().includes(text.toLowerCase()));
}

function includesAll(ary, otherAry) {
  return ary.length === intersection(ary, otherAry).length;
}

function arraysMatch(ary, otherAry, matchExact) {
  return (matchExact ? isEmpty(xor(ary, otherAry)) : !isEmpty(intersection(ary, otherAry)));
}

function isEmptyOrIncludes(aryOrStr, item) {
  return isEmpty(aryOrStr) || aryOrStr.includes(item);
}

function isEmptyOrIncludesAll(ary, otherAry) {
  return isEmpty(ary) || includesAll(ary, otherAry);
}

function isEmptyOrMatches(ary, otherAry, matchExact) {
  return isEmpty(ary) || arraysMatch(ary, otherAry, matchExact);
}

function isEmptyOrBaz(str, otherStr) {
  return isEmpty(str) || baz(str, otherStr);
}

function isAnyOrValue(a, b) {
  return ['Any', b].includes(a);
}

export {
  isEmptyOrIncludes,
  isEmptyOrMatches,
  isEmptyOrIncludesAll,
  isEmptyOrBaz,
  isAnyOrValue,
};
