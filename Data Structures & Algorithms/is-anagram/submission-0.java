class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;

        String s1 = s.toUpperCase();
        String t1 = t.toUpperCase();

        int [] counter = new int[26];

        for(int i = 0; i < s.length(); i++){
            counter[s1.charAt(i) - 'A'] ++ ; 
            counter[t1.charAt(i) - 'A'] -- ; 
        }
        for (int c : counter){
            if (c != 0) return false;
        }
        return true;
    }
}
