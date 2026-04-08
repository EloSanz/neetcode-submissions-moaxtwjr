class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();



        for (int n: nums){
            map.put(n, map.getOrDefault(n, 0) + 1 );
        }

        List<Map.Entry<Integer, Integer >> listaEntradas = new ArrayList<>(map.entrySet());
        listaEntradas.sort( (a,b) -> b.getValue() - a.getValue() );

        int [] resultado = new int[k];

        for(int i = 0; i < k; i++){
            resultado[i] = listaEntradas.get(i).getKey();
        }
        return resultado;
    }
}
