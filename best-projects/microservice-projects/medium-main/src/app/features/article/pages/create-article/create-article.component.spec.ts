import { ComponentFixture, TestBed } from '@angular/core/testing';
import CreateArticleComponent from '@article/pages/create-article/create-article.component';

describe('CreateArticleComponent', () => {
  let component: CreateArticleComponent;
  let fixture: ComponentFixture<CreateArticleComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [CreateArticleComponent],
    });
    fixture = TestBed.createComponent(CreateArticleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
