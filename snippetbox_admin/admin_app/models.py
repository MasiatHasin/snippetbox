from django.db import models
from django.utils import timezone

# Create your models here.
class Snippet(models.Model):
    id = models.AutoField(primary_key=True, editable=False)
    title = models.CharField(max_length=100)
    content = models.CharField(max_length=255)
    created_at = models.DateTimeField(default=timezone.now())
    expires_at = models.DateTimeField()
    
    class Meta:
        db_table = "snippets"

    def __str__(self) -> str:
        return f"{self.id}: {self.title} => {self.content[:100]}"
