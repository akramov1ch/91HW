# Uy vazifasi: Go va GitHub Actions yordamida RESTful API

## Maqsad
Bu vazifa `Go` da to‘liq ishlaydigan `RESTful API` yaratish, `GitHub Actions`, `Docker` va boshqa ishlab chiqarish mexanizmlarini avtomatlashtirish bilan ilg‘or `CI/CD` amaliyotlarini amalga oshirishni o‘z ichiga oladi.

## Talablar
1. **Go da RESTful API ishlab chiqish**:
    - Task Management API endpointlarni amalga oshiring (CRUD):

2. **Tasklarni saqlash uchun MongoDB o‘rnatish:**:
    - Vazifalarni quyidagi "field" bilan `MongoDB` da saqlang: `ID`, `Title`, `Description`, `Status (pending/completed)`, `CreatedAt`.
    - Endpointlar uchun mos xatolikni boshqarish va tekshirishni ta’minlang.
    
3. **Ko‘p bosqichli GitHub Actions workflowini yarating:**:
    - `1-bosqich: Testlash`:
        - go test yordamida unit testlarini bajaring.    
    - `2-bosqich: Build va Dockerize qiling`:
        - `Go` ilovasini build qiling va multi-stage Docker image yarating.
    - `3-bosqich: Docker imageni Docker Hub ga push qiling`:
 

4. **Kod va CI/CD pipeline sozlamalari bilan GitHub repository linkini yuboring.**














