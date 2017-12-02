//
//  AOCQ2Main.m
//  Q2
//
//  Created by nithin.g on 02/12/17.
//  Copyright © 2017 nithin.g. All rights reserved.
//

#import "AOCQ2Main.h"

@implementation AOCQ2Main

- (NSUInteger)calcChecksumFromString:(NSString *)spreadSheetContents{
    NSUInteger computedChecksum = 0;
    NSArray<NSString *> *rowEntriesList = [spreadSheetContents componentsSeparatedByString:@"\n"];
    for(NSString *rowEntry in rowEntriesList){
        NSRegularExpression *regex = [NSRegularExpression regularExpressionWithPattern:@"[^\\s]+" options:0 error:nil];
        
        __block NSMutableArray<NSNumber *> *rowEntriesList = [NSMutableArray new];
        [regex enumerateMatchesInString:rowEntry options:0 range:NSMakeRange(0, [rowEntry length]) usingBlock:^(NSTextCheckingResult * _Nullable result, NSMatchingFlags flags, BOOL * _Nonnull stop) {
            NSString *match = [rowEntry substringWithRange:[result range]];
            NSNumber *matchNum = [NSNumber numberWithInteger:[match integerValue]];
            [rowEntriesList addObject:matchNum];
        }];
        
        NSUInteger checksum = [self getChecksumForRow:[NSArray arrayWithArray:rowEntriesList]];
        computedChecksum += checksum;
        
    }
    return computedChecksum;
}

- (NSUInteger)getChecksumForRow:(NSArray<NSNumber *> *)rowEntriesList{
    NSUInteger checksum = 0;
    if([rowEntriesList count] == 0){
        return checksum;
    }
    
    // Part 1
    // checksum = [self getMaxMinDiffForRow:rowEntriesList];
    // Part 2
    checksum = [self getDivisibleQuotientForRow:rowEntriesList];
    return checksum;
}

- (NSUInteger)getMaxMinDiffForRow:(NSArray<NSNumber *> *)rowEntriesList{
    NSNumber *maxVal = [rowEntriesList valueForKeyPath:@"@max.intValue"];
    NSNumber *minVal = [rowEntriesList valueForKeyPath:@"@min.intValue"];
    return [maxVal unsignedIntegerValue] - [minVal unsignedIntegerValue];
}

- (NSUInteger)getDivisibleQuotientForRow:(NSArray<NSNumber *> *)rowEntriesList{
    NSArray<NSNumber *> *sortedList = [rowEntriesList sortedArrayUsingSelector:@selector(compare:)];
    for(int i=0; i<([sortedList count]-1);i++){
        NSUInteger smallNum = [sortedList[i] unsignedIntegerValue];
        for(int j=i+1; j<[sortedList count];j++){
            NSUInteger bigNum = [sortedList[j] unsignedIntegerValue];
            if((bigNum % smallNum) == 0){
                return bigNum/smallNum;
            }
        }
    }
    
    return 0;
}

@end
