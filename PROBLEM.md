## Problem: Encode Occurences

*EncodeOccurrences* adlı bir fonksiyon geliştirmeniz isteniyor. Bu fonksiyon verilen bir metin içerisinde geçen bazı özel kelimeleri belli bir kurala göre şifrelemeli ve metnin şifrelenmiş hali ekrana yazdırılmalı.

Fonksiyonun adı ve aldığı parametreler şu şekilde olmalı: (go)
`EncodeOccurrences(text string, words []string)`

**text:** Üzerinde işlem yapılacak metin

**words:** Metin içerisinde geçiyorsa şifrelenecek kelimelerin listesi

Metnin içinde bir özel kelimenin geçip geçmediğine bakılırken büyük-küçük harfe karşı duyarsız olunmalıdır.
```
words: "es", "sam"
text:  "theSE are sAMPles"
          ^^           ^^  ("es" 2 kere)
                  ^^^      ("sam" 1 kere)
```

**Şifreleme Kuralı:** Metin içerisindeki özel kelimelerle eşleşen yerlerde karakterlerin ötelenmesi, yani ingiliz alfabesinde belirtilen sayı kadar ileri götürülmesi gerekmektedir. Örneğin a harfi 1 defa ötelenirse b, 2 defa ötelenirse c olur. z'den sonra tekrar a geldiği kabul edilir.

Metinde tespit edilen bir özel kelime, metnin başından itibaren kaçıncı özel kelime ise o kadar öteleme yapılmalıdır. 

### Örnek: ###
```
Girdi: "pleAsE encode this", ["his", "ease"]
          ^^^^         ^^^
           1.           2.
e -> f
A -> B
s -> t
E -> F
(1 öteleme çünkü metinde bulunan 1. özel kelime)

h -> j
i -> k
s -> u
(2 öteleme çünkü metinde bulunan 2. özel kelime)


Çıktı: "plfBtF encode tjku"
```

### Kısıtlar: ###
- fmt kütüphanesi kullanımına izin verilmiyor. Onun yerine `utils` paketinde sadece tek bir karakter yazmamıza izin veren `PrintRune()` metodu kullanılabilecek.
- Herhangi bir hazır string metodu kullanımına izin verilmiyor. Bu problem için tamamen kendiniz stringler ve karakterler üzerinde döngüler kurarak çözüm yapmalısınız.