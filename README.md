# Productos API
login
```sh 
$ curl -v -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"name": "TeamBorregos","password":"1234"}'
``` 


https://www.amazon.com.mx/el%C3%A9ctrica-extensi%C3%B3n-Protector-Port%C3%A1tiles-dispositivos/dp/B01HPADPCW?pf_rd_r=K5ZF08VFY0HKEPK00HEW&pf_rd_p=e07157dc-8249-4cf7-8161-8070105962bf&pd_rd_r=94180c13-d95b-4f53-9f9f-be7f356398a8&pd_rd_w=ajNqK&pd_rd_wg=B3vUW&ref_=pd_gw_ci_mcx_mr_hp_d

https://www.amazon.com.mx/EletecPro-Tower-extensi%C3%B3n-protector-sobrecarga/dp/B074PMPM2B/ref=pd_sim_60_4/135-1383059-3977739?_encoding=UTF8&pd_rd_i=B074PMPM2B&pd_rd_r=9b69436b-9844-4e67-9025-2b384ab159e4&pd_rd_w=e1sJY&pd_rd_wg=fK2Jd&pf_rd_p=a62f455d-612d-4136-9fd7-44067fe2cd11&pf_rd_r=TNBFB7BPB1XWBHGF17YE&psc=1&refRID=TNBFB7BPB1XWBHGF17YE

https://www.amazon.com.mx/s?k=bonito+reloj+de+arena&__mk_es_MX=%C3%85M%C3%85%C5%BD%C3%95%C3%91&ref=nb_sb_noss

Opción 1 - REST API Style - cuando se selecciona un producto
backenddeedgar.com/search/reloj-arena
backenddeedgar.com/search/bulova
backenddeedgar.com/search/multicontacto
backenddeedgar.com/search/...
backenddeedgar.com/search/***cadena-a-buscar***   <- path param
backenddeedgar.com/search/:cadena-a-buscar   <- path param

Opción 2 - Query param - cuando se busca un producto
backenddeedgar.com/search?item=reloj-arena&

https://www.google.com/aclk?sa=L&ai=DChcSEwitrObC7dDqAhUB28AKHRrVCMcYABABGgJpbQ&sig=AOD64_07bUda2OZbrSEYZCEPpJf5TvJojA&q=&ved=2ahUKEwiaot_C7dDqAhVMbKwKHaGpDUsQ0Qx6BAgcEAE&adurl=

https://intl.bulova.com/?gclid=EAIaIQobChMIrazmwu3Q6gIVAdvACh0a1QjHEAAYASAAEgKTmPD_BwE

https://www.linio.com.mx/?adjust_t=1zira0_f1h7ws&adjust_google_network=g&adjust_google_placement=&adjust_campaign=mex-sembr-sear_Linio&adjust_adgroup=92121284556&gclid=EAIaIQobChMIuKLJwO7Q6gIVEdvACh1gWwGfEAAYASAAEgIsE_D_BwE

https://www.linio.com.mx/search?scroll=&q=bulova

https://www.linio.com.mx/p/reloj-bulova-crono-grafo-precisionist-uhf-98b315-time-squaretm-lfnlkw?qid=a5314909f6eeb5748a4ee2e3af2482ef&oid=BU245FA1DG38YLMX&position=55&sku=BU245FA1DG38YLMX

how to make a call
```
curl -i http://127.0.0.1:3030/search\?q\=cadena
```
const cadena-hija = "hija de la hija"
`cadena padre ${cadena-hija}`
"cadena padre hija de la hija"

cadena-hija := "hija de la hija"
fmt.Sprintf("cadena padre %s", cadena-hija)

AMZ_BASE_URL = https://api.rainforestapi.com/request?api_key=2AFBB75C890447B4B31D97B70028C9DA&type=search&amazon_domain=amazon.com.mx&search_term=

https://api.rainforestapi.com/request?api_key=2AFBB75C890447B4B31D97B70028C9DA&type=search&amazon_domain=amazon.com&search_term=%s&associate_id=mejorprecio