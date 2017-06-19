
# FOD
![alt text](https://backtory.com/landingpageComponents/images/H40pxB.png)

Freedom of Developers
Please put domains of websites and web apps you need as a developer that somehow! you cant access directly  in "domains" file , then use this proxy :
```
fod.backtory.com:8118
```

* FOD.ppx
Import this file in proxifire application (https://www.proxifier.com/) (Win/Mac) to transparent your TCP connection from any application trough fod.backtory.com proxy.
All other addresses that are not defined in "domains" file will route directly.

<h1 lang="fa" dir="rtl" align="right">توسعه نرم‌افزار بدون زجر</h1>
```
*.developer.android.com
*.atlassian.com
*.bintray.com
*.bitbucket.org
*.developers.google.com
*.domains.google.com
*.dl.google.com
*.appengine.google.com
*.google.ai
*.unity3d.com
*.cloud.google.com
*.analytics.google.com
*.fiber.google.com
*.dl-ssl.google.com
*.dns.google.com
*.firebase.google.com
*.firebase.com
*.googleapis.com
*.java.com
*.khanacademy.org
*.oracle.com
*.storage.googleapis.com
*.download.virtualbox.org
*.hub.docker.com
*.mongodb.org
*.adobe.com
*.vmware.com
*.toggl.com
*.cloud.docker.com
*.googlesource.com
*.polymer-project.org
*.material.io
*.teamviewer.com
*.intel.com
.developer.chrome.com
```
<p lang="fa" dir="rtl" align="right">سرورهای بکتوری در دیتاسنترهای داخل کشور قرار دارد و در نتیجه با توجه به شرایط اینترنت کشور شما می‌توانید با سرعت بیشتری دیتا مورد نیاز خود را دانلود کنید. اگر اطلاعات کافی برای استفاده از پراکسی ندارید سعی کردیم آموزش‌های مختصری برای شما فراهم کنیم. سعی ما بر این است به مرور زمان سرویس و آموزش‌ها را بهبود دهیم تا نیازهای بیشتری را برطرف کند.</p>
<p lang="fa" dir="rtl" align="right">توجه داشته باشید که این سرویس فقط برای دور زدن تحریم‌ها ساخته شده و سایت‌هایی که تحریم نکردند یا فیلتر شدن از این سرویس قابل دسترسی نیست.</p>
<h2 lang="fa" dir="rtl" align="right">اطلاعات سرویس</h2>
<p lang="fa" dir="rtl" align="right">برای استفاده از پراکسی می‌توانید از این آدرس و پورت استفاده کنید</p>
```
address: fod.backtory.com
port:8118
```
<p lang="fa" dir="rtl" align="right">کافیست اطلاعات بالا در تنظیمات پراکسی برنامه مورد نظر خود وارد کنید.</p>
<h2 lang="fa" dir="rtl" align="right">افزودن به سرویس</h2>
<p lang="fa" dir="rtl" align="right">اگر از سایت یا سرویسی استفاده می‌کنید که ایران تحریم کرده ولی در این سرویس نیست می‌توانید با زدن ایشو یا افزودن دامین مورد نظر در فایل domains همین ریپو و پول ریکوست به ما اعلام کنید تا اضافش کنیم.</p>
<h2 lang="fa" dir="rtl" align="right">Android Studio</h2>
<p lang="fa" dir="rtl" align="right">وارد setting شوید، در لینوکس از منوی file (ویندوز و مک هم شبیه هستند)<br>در کادر جست‌جو عبارت proxy وارد کنید تا صفحه http proxy آورده شود<br>گزینه Manual proxy configuration انتخاب کنید سپس http<br>در مقابل Host name آدرس پراکسی و در مقابل Port number پورت گفته شده وارد کنید</p>
<h2 lang="fa" dir="rtl" align="right">Gradle</h2>
<p lang="fa" dir="rtl" align="right">در فایل gradle.properties خط‌های زیر را به همراه اطلاعات پراکسی اضافه کنید</p>
```
systemProp.http.proxyHost=fod.backtory.com
systemProp.http.proxyPort=8118
```
<p lang="fa" dir="rtl" align="right">در فایل توجه داشته باشید که اگر از ریپازیتوری‌ای جز jcenter استفاده کنید احتمال داره به مشکل بخورید.<br>در صورت امکان از ترنسپرنت پراکسی استفاده کنید و فقط دامین‌های موجود رو از پراکسی رد کنید.<br> در صورتی که با gradle به مشکل خوردید به ما بگید تا دنبال راه حل بهتری باشیم.</p>
<h2 lang="fa" dir="rtl" align="right">فایرفاکس</h2>
<p lang="fa" dir="rtl" align="right"><a href="https://addons.mozilla.org/en-US/firefox/addon/foxyproxy-standard/">https://addons.mozilla.org/en-US/firefox/addon/foxyproxy-standard</a><br>foxyproxy یه پلاگین برای فایرفکس و  فکر کنم کروم هست که می‌تونید خیلی راحت و دم دستی تنظیمات پراکسی رو توش تغییر بدید.<br>پلاگین رو نصب کنید، add new proxy رو بزنید، تنظیمات پراکسی رو وارد کنید و ذخیره کنید.<br>از قسمت مود پراکسی‌ای که ساختید رو فعال کنید.</p>
<p lang="fa" dir="rtl" align="right">نکته:‌بدون این پلاگین هم میشه از تنظیمات فایرفکس پراکسی رو تغییر داد ولی این دم دسته</p>
<p lang="fa" dir="rtl" align="right">می‌تونید تنظیمات دیگه‌ای هم بهش اضافه کنید تا فقط سایت‌های موجود در سرویس از پراکسی رد شند باشه</p>

<h2 lang="fa" dir="rtl" align="right">Chrome</h2>
<p lang="fa" dir="rtl" align="right"><a href="https://github.com/FelisCatus/SwitchyOmega/releases">https://github.com/FelisCatus/SwitchyOmega/releases</a><br>اکستنشنی هست به نام SwitchyOmega که می‌تونه تنظیمات پراکسی کروم رو باهاش دست کاری کرد. نصبش کنید.<br>کنار آدرس بار آیکنش اضافه می‌شه از option گزینه new profile رو انتخاب کنید، یه اسم براش وارد کنید و گزینه proxy profile رو بزنید. بعد از وارد کردن اطلاعات پراکسی apply change رو بزنید.<br>هر وقت خواستید می‌تونید با کلیک رو آیکنش به راحتی ارتباط رو مستقیم کنید یا از پراکسی رد کنید.<br>تنظیماتی داره که چه سایت‌هایی از چه پراکسی‌ای رد شند می‌تونید بگید سایت‌هایی که در این سوریس وجود دارند از این پراکسی رد شند و بقیه سایت‌ها به روش دیگری</p>
<h2 lang="fa" dir="rtl" align="right">ProxyFire</h2>
<p lang="fa" dir="rtl" align="right">با پراکسی فایر در لینوکس و مک یه جورایی میشه یه ترنسپرنت پراکسی راه انداخت و بدون اینکه برای بقیه برنامه‌ها تنظیماتی انجام داد ترافیکشون رو از پراکسی رد کنید همچنین اگر با gradle به مشکل خوردید احتمالا این روش جواب گو خواهد بود<br>با یک جستجو در اینترنت  یه نسخه از برنامه رو دانلود کنید و سپس فایل fox.ppx در همین ریپازیتوری را دانلود و در برنامه ایمپورت کنید.</p>
<p lang="fa" dir="rtl" align="right">ادامه دارد</p>
<p><a href="http://creativecommons.org/licenses/by-sa/3.0/">http://creativecommons.org/licenses/by-sa/3.0</a>

