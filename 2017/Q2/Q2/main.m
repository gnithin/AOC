//
//  main.m
//  Q2
//
//  Creatxed by nithin.g on 02/12/17.
//  Copyright © 2017 nithin.g. All rights reserved.
//

#import <Foundation/Foundation.h>
#import "AOCQ2Main.h"

void customMain(void);

int main(int argc, const char * argv[]) {
    @autoreleasepool {
        customMain();
    }
    return 0;
}

void customMain(void){
    NSString *filePath = [[NSBundle mainBundle] pathForResource:@"ip" ofType:@"txt"];
    
    NSError *err;
    NSString *fileContents = [NSString stringWithContentsOfFile:filePath encoding:NSUTF8StringEncoding error:&err];
    if(err){
        NSLog(@"Reading file returned an error - %@", err);
        return;
    }
    AOCQ2Main *main = [AOCQ2Main new];
    NSUInteger checksum = [main calcChecksumFromString:fileContents];
    NSLog(@"The checksum is %lu", checksum);
}
