func romanToInt(s string) int {
    
    rMapper := map[rune]int{
        'I' : 1,
        'V' : 5,
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }
    
    
    var res int
    
    
    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && rMapper[rune(s[i])] < rMapper[rune(s[i+1])]{
            
            res -= rMapper[rune(s[i])]
            
            continue
            
        }
        res += rMapper[rune(s[i])] 
        
    }
    
    return res
    
}