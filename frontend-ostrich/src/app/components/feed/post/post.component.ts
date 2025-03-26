import { Component, Input, OnInit } from '@angular/core';
import { FeedData } from '../../../services/models/Feed.Module';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-post',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './post.component.html',
  styleUrl: './post.component.css'
})
export class PostComponent implements OnInit {
  @Input() post: FeedData = {
    authorPubKey: 'teste',
    authorName: 'teste',
    authorImage: 'teste',
    content: 'content teste',
    tags: [],
    timestamp: 1743005638
  }

  public images:     string[] = [];
  public videos:     string[] = [];
  public links:      string[] = [];
  public hashtags:   string[] = [];
  public showImages: boolean  = false;
  public showVideos: boolean  = false;
  public showLinks:  boolean  = false;
  public showTags:   boolean  = false;
  public isContentExpanded: boolean = false;

  constructor() { }

  ngOnInit(): void {
    this.extractLinksAndTags(this.post.content);
  }

  formatDate(timestamp: number): string {
    const date = new Date(timestamp * 1000);
    return date.toLocaleDateString("pt-BR", {
      day: "2-digit",
      month: "short",
      year: "numeric",
    }) + " " + date.toLocaleTimeString("pt-BR", {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit"
    });
  }

  toggleContent(): void {
    this.isContentExpanded = !this.isContentExpanded;
  }

  toggleSection(section: string): void {
    if (section === 'images') {
      this.showImages = !this.showImages;
    } else if (section === 'videos') {
      this.showVideos = !this.showVideos;
    } else if (section === 'links') {
      this.showLinks = !this.showLinks;
    } else if (section === 'tags') {
      this.showTags = !this.showTags;
    }
  }

  // Função para identificar links, imagens, vídeos e tags
  extractLinksAndTags(content: string): void {
    const imageRegex = /\bhttps?:\/\/(?:\S+\.(?:jpg|jpeg|png|gif|bmp|svg))\b/gi; // Expressão para links de imagem
    const videoRegex = /\bhttps?:\/\/(?:www\.)?(?:youtube\.com\/(?:watch\?v=|v\/)|vimeo\.com\/)([a-zA-Z0-9_-]{11})\b/gi; // Expressão para links de vídeo (YouTube/Vimeo)
    const linkRegex = /\bhttps?:\/\/\S+/gi; // Expressão para qualquer link genérico
    const tagRegex = /\#\w+/g; // Expressão para identificar tags (links que começam com '#')

    // Encontrar todos os links
    const images = content.match(imageRegex);
    const videos = content.match(videoRegex);
    const links = content.match(linkRegex);
    const tags = content.match(tagRegex);

    if (images) {
      this.images = images;
    }

    if (videos) {
      this.videos = videos;
    }

    if (links) {
      this.links = links.filter(link => !this.images.includes(link) && !this.videos.includes(link));
    }

    if (tags) {
      this.hashtags = tags;
    }
  }

  getVideoEmbedUrl(videoUrl: string): string {
    const youtubeRegex = /(?:youtube\.com\/(?:watch\?v=|v\/)|youtu\.be\/)([a-zA-Z0-9_-]{11})/i;
    const vimeoRegex = /(?:vimeo\.com\/)([0-9]+)/i;
  
    let match = videoUrl.match(youtubeRegex);
    if (match) {
      return `https://www.youtube.com/embed/${match[1]}`;
    }
  
    match = videoUrl.match(vimeoRegex);
    if (match) {
      return `https://player.vimeo.com/video/${match[1]}`;
    }
  
    return videoUrl; // Retorna o link diretamente caso não seja YouTube ou Vimeo
  }
}
