import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input } from '@angular/core';

const TagListImports: Array<any> = [CommonModule];

@Component({
  selector: 'app-tag-list',
  standalone: true,
  imports: TagListImports,
  template: `
    <ul class="tag-list">
      <li *ngFor="let tag of tagList" class="tag-list__item">
        {{ tag }}
      </li>
    </ul>
  `,
  styleUrls: ['./tag-list.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TagListComponent {
  @Input() public tagList: string[] = ['dwa', 'testttt', 'tes', 'testtttttt'];
}
