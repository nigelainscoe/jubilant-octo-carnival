# Claude Code Skills Manager - Product Specification

**Version:** 1.0.0
**Date:** December 2024
**Status:** Draft

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Product Overview](#2-product-overview)
3. [User Stories](#3-user-stories)
4. [Technical Architecture](#4-technical-architecture)
5. [Database Schema](#5-database-schema)
6. [API Design](#6-api-design)
7. [UI/UX Design](#7-uiux-design)
8. [LLM Integration](#8-llm-integration)
9. [Skill Export Format](#9-skill-export-format)
10. [Security Considerations](#10-security-considerations)
11. [Implementation Phases](#11-implementation-phases)
12. [Appendices](#12-appendices)

---

## 1. Executive Summary

### 1.1 Purpose

The Claude Code Skills Manager is a local web application designed to streamline the creation, management, and export of Claude Code Skills. It provides an intuitive interface for users to describe their skill requirements, collaborates with a Claude LLM to iteratively develop comprehensive skill packages, and stores completed skills in a local database for future reference and export.

### 1.2 Key Features

- **Interactive Skill Creation**: Describe skills in natural language with iterative refinement through Claude LLM Q&A
- **AI-Powered Generation**: Automatic generation of `SKILL.md` files, assets, and references
- **Persistent Storage**: DuckDB-backed local database for skill management
- **Export Functionality**: Package skills as `.skill` compressed archives
- **Minecraft-Themed UI**: Distinctive visual experience using a pixel-art, block-based design aesthetic

### 1.3 Technology Stack

| Layer | Technology |
|-------|------------|
| Frontend | TypeScript, TailwindCSS, Vanilla JS or React |
| Backend | Node.js with TypeScript |
| Database | DuckDB |
| LLM | Claude API (Anthropic) |
| Build Tools | Vite or similar bundler |

---

## 2. Product Overview

### 2.1 Problem Statement

Creating Claude Code Skills manually requires:
- Understanding the SKILL.md format and YAML frontmatter requirements
- Proper directory structure organization (scripts/, references/, assets/)
- Knowledge of progressive disclosure architecture
- Testing and iteration to refine skill behavior

This process is time-consuming and error-prone, especially for users unfamiliar with skill authoring best practices.

### 2.2 Solution

A guided, AI-assisted skill creation tool that:
1. Accepts natural language descriptions of desired skill functionality
2. Engages in clarifying dialogue to refine requirements
3. Generates properly structured skill packages
4. Validates output against Claude Code skill specifications
5. Stores and manages skills locally
6. Enables easy export for deployment

### 2.3 Target Users

- Developers building custom Claude Code workflows
- Teams standardizing Claude Code usage patterns
- Power users creating personal productivity skills
- Organizations distributing internal tooling via skills

### 2.4 Out of Scope (v1.0)

- Cloud synchronization of skills
- Multi-user collaboration features
- Skill marketplace/sharing platform
- Direct deployment to Claude Code
- Version control integration

---

## 3. User Stories

### 3.1 Skill Creation

**US-001**: As a user, I want to describe what I need a skill to do in plain English so that I don't need to know the technical skill format.

**US-002**: As a user, I want the AI to ask me clarifying questions so that the generated skill accurately reflects my needs.

**US-003**: As a user, I want to see a preview of the generated SKILL.md before finalizing so that I can request changes.

**US-004**: As a user, I want the AI to suggest assets and references that would enhance my skill so that I create comprehensive skill packages.

**US-005**: As a user, I want to iterate on a skill through conversation so that I can refine it to perfection.

### 3.2 Skill Management

**US-006**: As a user, I want to view all my saved skills in a list so that I can manage my skill library.

**US-007**: As a user, I want to search and filter my skills so that I can quickly find what I need.

**US-008**: As a user, I want to edit an existing skill so that I can update it as requirements change.

**US-009**: As a user, I want to delete skills I no longer need so that I can keep my library organized.

**US-010**: As a user, I want to duplicate a skill so that I can create variations without starting from scratch.

### 3.3 Skill Export

**US-011**: As a user, I want to export a skill as a `.skill` archive so that I can deploy it to Claude Code.

**US-012**: As a user, I want to export multiple skills at once so that I can batch deploy my skill library.

**US-013**: As a user, I want to import a `.skill` archive so that I can manage externally-created skills.

### 3.4 Configuration

**US-014**: As a user, I want to configure my Claude API key so that the application can generate skills.

**US-015**: As a user, I want to set default skill metadata (author, license) so that I don't repeat this information.

---

## 4. Technical Architecture

### 4.1 High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                        Browser (Frontend)                        │
│  ┌───────────────┐  ┌───────────────┐  ┌───────────────────┐   │
│  │   Skill List  │  │ Skill Editor  │  │ Conversation View │   │
│  │     View      │  │     View      │  │                   │   │
│  └───────────────┘  └───────────────┘  └───────────────────┘   │
│                              │                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │              TailwindCSS + Minecraft Theme               │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
                               │
                          HTTP/REST
                               │
┌─────────────────────────────────────────────────────────────────┐
│                     Node.js Backend (Express)                    │
│  ┌───────────────┐  ┌───────────────┐  ┌───────────────────┐   │
│  │  Skills API   │  │   LLM API     │  │   Export API      │   │
│  │  Controller   │  │  Controller   │  │   Controller      │   │
│  └───────────────┘  └───────────────┘  └───────────────────┘   │
│                              │                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │                    Service Layer                          │   │
│  │  ┌────────────┐  ┌────────────┐  ┌──────────────────┐   │   │
│  │  │SkillService│  │ LLMService │  │ConversationService│   │   │
│  │  └────────────┘  └────────────┘  └──────────────────┘   │   │
│  └──────────────────────────────────────────────────────────┘   │
│                              │                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │                   Data Access Layer                       │   │
│  │              DuckDB Repository Classes                    │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                         DuckDB Database                          │
│  ┌───────────┐  ┌────────────┐  ┌─────────────┐  ┌──────────┐  │
│  │  skills   │  │   assets   │  │ references  │  │ messages │  │
│  └───────────┘  └────────────┘  └─────────────┘  └──────────┘  │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                        Claude API (External)                     │
└─────────────────────────────────────────────────────────────────┘
```

### 4.2 Directory Structure

```
skills-manager/
├── src/
│   ├── client/                    # Frontend code
│   │   ├── components/            # UI components
│   │   │   ├── SkillList.ts
│   │   │   ├── SkillEditor.ts
│   │   │   ├── ConversationView.ts
│   │   │   ├── ExportModal.ts
│   │   │   └── common/
│   │   │       ├── MinecraftButton.ts
│   │   │       ├── MinecraftCard.ts
│   │   │       ├── MinecraftInput.ts
│   │   │       └── MinecraftModal.ts
│   │   ├── styles/
│   │   │   ├── main.css           # TailwindCSS entry
│   │   │   └── minecraft-theme.css # Custom Minecraft styles
│   │   ├── services/
│   │   │   └── api.ts             # API client
│   │   ├── types/
│   │   │   └── index.ts           # Shared types
│   │   └── main.ts                # Entry point
│   │
│   ├── server/                    # Backend code
│   │   ├── controllers/
│   │   │   ├── skillsController.ts
│   │   │   ├── llmController.ts
│   │   │   └── exportController.ts
│   │   ├── services/
│   │   │   ├── skillService.ts
│   │   │   ├── llmService.ts
│   │   │   ├── conversationService.ts
│   │   │   └── exportService.ts
│   │   ├── repositories/
│   │   │   ├── skillRepository.ts
│   │   │   ├── assetRepository.ts
│   │   │   ├── referenceRepository.ts
│   │   │   └── messageRepository.ts
│   │   ├── models/
│   │   │   └── index.ts           # Database models
│   │   ├── middleware/
│   │   │   └── errorHandler.ts
│   │   ├── database/
│   │   │   ├── connection.ts
│   │   │   └── migrations/
│   │   ├── prompts/
│   │   │   └── skillGeneration.ts # LLM prompts
│   │   └── index.ts               # Server entry
│   │
│   └── shared/                    # Shared code
│       ├── types/
│       │   └── index.ts
│       └── validators/
│           └── skillValidator.ts
│
├── public/
│   ├── fonts/                     # Minecraft-style fonts
│   └── images/                    # UI assets
│
├── data/                          # DuckDB database files
│   └── skills.db
│
├── package.json
├── tsconfig.json
├── tailwind.config.js
├── vite.config.ts
└── README.md
```

### 4.3 Component Responsibilities

#### 4.3.1 Frontend Components

| Component | Responsibility |
|-----------|----------------|
| `SkillList` | Display paginated list of skills with search/filter |
| `SkillEditor` | View/edit skill details, preview SKILL.md |
| `ConversationView` | Chat interface for skill creation/refinement |
| `ExportModal` | Configure and execute skill exports |
| `MinecraftButton` | Styled button component with Minecraft aesthetics |
| `MinecraftCard` | Styled card component for skill display |
| `MinecraftInput` | Styled form inputs |
| `MinecraftModal` | Styled modal dialogs |

#### 4.3.2 Backend Services

| Service | Responsibility |
|---------|----------------|
| `SkillService` | CRUD operations for skills, validation |
| `LLMService` | Claude API interaction, prompt management |
| `ConversationService` | Manage conversation history, context |
| `ExportService` | Generate .skill archives |

---

## 5. Database Schema

### 5.1 DuckDB Schema Definition

```sql
-- Skills table: Core skill metadata and content
CREATE TABLE skills (
    id VARCHAR PRIMARY KEY,              -- UUID
    name VARCHAR(64) NOT NULL,           -- Skill name (lowercase, hyphens)
    description VARCHAR(1024) NOT NULL,  -- Brief description
    skill_md_content TEXT NOT NULL,      -- Full SKILL.md content
    allowed_tools VARCHAR,               -- Comma-separated tool restrictions
    model VARCHAR,                       -- Target Claude model
    version VARCHAR DEFAULT '1.0.0',     -- Semantic version
    license VARCHAR,                     -- License identifier
    author VARCHAR,                      -- Author name
    tags VARCHAR,                        -- Comma-separated tags
    status VARCHAR DEFAULT 'draft',      -- draft, complete, archived
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Assets table: Binary files and templates
CREATE TABLE assets (
    id VARCHAR PRIMARY KEY,              -- UUID
    skill_id VARCHAR NOT NULL,           -- FK to skills
    filename VARCHAR NOT NULL,           -- Original filename
    file_path VARCHAR NOT NULL,          -- Path within skill directory
    content_type VARCHAR NOT NULL,       -- MIME type
    content BLOB NOT NULL,               -- Binary content
    size_bytes INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);

-- References table: Documentation files
CREATE TABLE skill_references (
    id VARCHAR PRIMARY KEY,              -- UUID
    skill_id VARCHAR NOT NULL,           -- FK to skills
    filename VARCHAR NOT NULL,           -- e.g., 'schemas.md'
    file_path VARCHAR NOT NULL,          -- Path within skill directory
    content TEXT NOT NULL,               -- Markdown content
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);

-- Scripts table: Executable script files
CREATE TABLE scripts (
    id VARCHAR PRIMARY KEY,              -- UUID
    skill_id VARCHAR NOT NULL,           -- FK to skills
    filename VARCHAR NOT NULL,           -- e.g., 'process.py'
    file_path VARCHAR NOT NULL,          -- Path within skill directory
    content TEXT NOT NULL,               -- Script content
    language VARCHAR NOT NULL,           -- python, bash, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);

-- Conversations table: Skill creation conversations
CREATE TABLE conversations (
    id VARCHAR PRIMARY KEY,              -- UUID
    skill_id VARCHAR,                    -- FK to skills (nullable until saved)
    title VARCHAR NOT NULL,              -- Conversation title
    status VARCHAR DEFAULT 'active',     -- active, completed, abandoned
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE SET NULL
);

-- Messages table: Individual conversation messages
CREATE TABLE messages (
    id VARCHAR PRIMARY KEY,              -- UUID
    conversation_id VARCHAR NOT NULL,    -- FK to conversations
    role VARCHAR NOT NULL,               -- 'user' or 'assistant'
    content TEXT NOT NULL,               -- Message content
    message_type VARCHAR DEFAULT 'text', -- text, skill_preview, question
    metadata JSON,                       -- Additional structured data
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);

-- Settings table: Application configuration
CREATE TABLE settings (
    key VARCHAR PRIMARY KEY,
    value TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for common queries
CREATE INDEX idx_skills_status ON skills(status);
CREATE INDEX idx_skills_name ON skills(name);
CREATE INDEX idx_skills_created ON skills(created_at);
CREATE INDEX idx_assets_skill ON assets(skill_id);
CREATE INDEX idx_references_skill ON skill_references(skill_id);
CREATE INDEX idx_scripts_skill ON scripts(skill_id);
CREATE INDEX idx_messages_conversation ON messages(conversation_id);
CREATE INDEX idx_conversations_skill ON conversations(skill_id);
```

### 5.2 Entity Relationships

```
┌─────────────┐      1:N      ┌─────────────┐
│   skills    │──────────────▶│   assets    │
└─────────────┘               └─────────────┘
       │
       │      1:N      ┌─────────────────┐
       │──────────────▶│ skill_references│
       │               └─────────────────┘
       │
       │      1:N      ┌─────────────┐
       │──────────────▶│   scripts   │
       │               └─────────────┘
       │
       │      1:N      ┌───────────────┐      1:N      ┌──────────┐
       └──────────────▶│ conversations │──────────────▶│ messages │
                       └───────────────┘               └──────────┘
```

---

## 6. API Design

### 6.1 REST Endpoints

#### 6.1.1 Skills API

```
GET    /api/skills                  # List all skills (with pagination, search)
GET    /api/skills/:id              # Get single skill with all components
POST   /api/skills                  # Create new skill
PUT    /api/skills/:id              # Update skill
DELETE /api/skills/:id              # Delete skill
POST   /api/skills/:id/duplicate    # Duplicate a skill
```

#### 6.1.2 Assets API

```
GET    /api/skills/:skillId/assets           # List skill assets
POST   /api/skills/:skillId/assets           # Upload asset
GET    /api/skills/:skillId/assets/:id       # Download asset
DELETE /api/skills/:skillId/assets/:id       # Delete asset
```

#### 6.1.3 References API

```
GET    /api/skills/:skillId/references       # List skill references
POST   /api/skills/:skillId/references       # Create reference
PUT    /api/skills/:skillId/references/:id   # Update reference
DELETE /api/skills/:skillId/references/:id   # Delete reference
```

#### 6.1.4 Scripts API

```
GET    /api/skills/:skillId/scripts          # List skill scripts
POST   /api/skills/:skillId/scripts          # Create script
PUT    /api/skills/:skillId/scripts/:id      # Update script
DELETE /api/skills/:skillId/scripts/:id      # Delete script
```

#### 6.1.5 Conversations API

```
GET    /api/conversations                    # List conversations
GET    /api/conversations/:id                # Get conversation with messages
POST   /api/conversations                    # Start new conversation
POST   /api/conversations/:id/messages       # Send message (triggers LLM)
DELETE /api/conversations/:id                # Delete conversation
POST   /api/conversations/:id/complete       # Mark complete & save skill
```

#### 6.1.6 Export API

```
POST   /api/export/skill/:id                 # Export single skill
POST   /api/export/skills                    # Export multiple skills
POST   /api/import/skill                     # Import .skill archive
```

#### 6.1.7 Settings API

```
GET    /api/settings                         # Get all settings
PUT    /api/settings                         # Update settings
```

### 6.2 Request/Response Examples

#### Create Skill (POST /api/skills)

**Request:**
```json
{
  "name": "sql-query-helper",
  "description": "Helps write and optimize SQL queries for PostgreSQL databases",
  "skill_md_content": "---\nname: sql-query-helper\ndescription: Helps write and optimize SQL queries\n---\n\n# SQL Query Helper\n\n...",
  "allowed_tools": "Read,Grep,Bash",
  "tags": "sql,database,postgresql"
}
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "sql-query-helper",
  "description": "Helps write and optimize SQL queries for PostgreSQL databases",
  "skill_md_content": "...",
  "allowed_tools": "Read,Grep,Bash",
  "model": null,
  "version": "1.0.0",
  "license": null,
  "author": null,
  "tags": "sql,database,postgresql",
  "status": "draft",
  "created_at": "2024-12-15T10:30:00Z",
  "updated_at": "2024-12-15T10:30:00Z"
}
```

#### Send Conversation Message (POST /api/conversations/:id/messages)

**Request:**
```json
{
  "content": "I want a skill that helps me write unit tests for React components using Jest and React Testing Library"
}
```

**Response (streamed or complete):**
```json
{
  "message": {
    "id": "msg-123",
    "role": "assistant",
    "content": "I'd be happy to help you create a React testing skill! Before I generate the skill, I have a few clarifying questions:\n\n1. **Component Types**: Are you primarily testing functional components with hooks, or do you also need support for class components?\n\n2. **Testing Patterns**: Which testing patterns do you want to emphasize?\n   - User interaction testing (clicks, form inputs)\n   - Snapshot testing\n   - Mock/spy patterns for API calls\n   - Accessibility testing\n\n3. **State Management**: Do your components use any state management libraries (Redux, Zustand, etc.) that should be considered?\n\n4. **Custom Hooks**: Should the skill include guidance for testing custom hooks with `@testing-library/react-hooks`?\n\nPlease let me know your preferences, and I'll create a comprehensive skill tailored to your needs.",
    "message_type": "question",
    "created_at": "2024-12-15T10:32:00Z"
  }
}
```

#### Export Skill (POST /api/export/skill/:id)

**Request:**
```json
{
  "include_conversation": false
}
```

**Response:**
```
Content-Type: application/zip
Content-Disposition: attachment; filename="sql-query-helper.skill"

[Binary ZIP data]
```

---

## 7. UI/UX Design

### 7.1 Design Philosophy

The application adopts a Minecraft-inspired aesthetic that:
- Uses pixelated, block-style visual elements
- Employs the Minecraft color palette (dirt browns, grass greens, stone grays)
- Features pixel art iconography
- Includes subtle animations reminiscent of Minecraft UI interactions
- Maintains readability and usability despite thematic styling

### 7.2 Color Palette

```css
/* Primary Colors */
--mc-dirt-brown: #866043;
--mc-grass-green: #5B8731;
--mc-stone-gray: #7F7F7F;
--mc-wood-tan: #9C7E4E;
--mc-water-blue: #2859A5;

/* UI Colors */
--mc-ui-dark: #1D1D1D;
--mc-ui-medium: #373737;
--mc-ui-light: #C6C6C6;
--mc-ui-highlight: #5A5A5A;

/* Status Colors */
--mc-success: #47A025;      /* Emerald green */
--mc-warning: #DAA520;      /* Gold/amber */
--mc-error: #B22222;        /* Redstone red */
--mc-info: #4169E1;         /* Lapis blue */

/* Text Colors */
--mc-text-primary: #FFFFFF;
--mc-text-secondary: #AAAAAA;
--mc-text-shadow: #3F3F3F;
```

### 7.3 Typography

Use a pixel-style font for headings and UI elements:
- **Primary Font**: "MinecraftSeven" or similar pixel font
- **Secondary Font**: System monospace for code blocks
- **Body Text**: Sans-serif with appropriate sizing for readability

```css
/* Tailwind font configuration */
fontFamily: {
  'minecraft': ['MinecraftSeven', 'VT323', 'Press Start 2P', 'monospace'],
  'body': ['Inter', 'system-ui', 'sans-serif'],
  'code': ['Fira Code', 'Monaco', 'Consolas', 'monospace']
}
```

### 7.4 Component Styling

#### 7.4.1 Minecraft Button

```css
/* Base button style - pixelated 3D effect */
.mc-button {
  @apply relative px-6 py-3 font-minecraft text-white uppercase;
  @apply bg-[#7F7F7F] border-4;
  border-style: solid;
  border-color: #FFFFFF #555555 #555555 #FFFFFF;
  image-rendering: pixelated;
  text-shadow: 2px 2px 0 #3F3F3F;
}

.mc-button:hover {
  @apply bg-[#A0A0A0];
  border-color: #FFFFFF #666666 #666666 #FFFFFF;
}

.mc-button:active {
  @apply bg-[#6A6A6A];
  border-color: #555555 #FFFFFF #FFFFFF #555555;
}

/* Primary variant - grass green */
.mc-button-primary {
  @apply bg-[#5B8731];
  border-color: #8BC34A #2E5016 #2E5016 #8BC34A;
}
```

#### 7.4.2 Minecraft Card

```css
.mc-card {
  @apply relative p-4;
  background: linear-gradient(180deg, #373737 0%, #1D1D1D 100%);
  border: 4px solid;
  border-color: #5A5A5A #0A0A0A #0A0A0A #5A5A5A;
  box-shadow:
    inset 2px 2px 0 rgba(255,255,255,0.1),
    inset -2px -2px 0 rgba(0,0,0,0.3);
}

.mc-card-header {
  @apply font-minecraft text-lg text-white mb-2;
  text-shadow: 2px 2px 0 #3F3F3F;
}
```

#### 7.4.3 Minecraft Input

```css
.mc-input {
  @apply w-full px-4 py-2 font-body text-white;
  background: #000000;
  border: 3px solid;
  border-color: #373737 #5A5A5A #5A5A5A #373737;
}

.mc-input:focus {
  @apply outline-none;
  border-color: #5B8731 #2E5016 #2E5016 #5B8731;
  box-shadow: 0 0 0 2px rgba(91, 135, 49, 0.3);
}
```

### 7.5 Page Layouts

#### 7.5.1 Main Layout

```
┌──────────────────────────────────────────────────────────────────────┐
│  ████  SKILLS MANAGER  ████                        [Settings] [Help] │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │
├──────────────────────────────────────────────────────────────────────┤
│                                                                       │
│  [+ New Skill]        [Search: ________________] [Filter ▼] [Export] │
│                                                                       │
│  ┌─────────────────────────────────────────────────────────────────┐ │
│  │ ⬜ sql-query-helper                              [Edit] [Delete] │ │
│  │    Helps write and optimize SQL queries                         │ │
│  │    Status: Complete  |  Tags: sql, database  |  v1.0.0          │ │
│  ├─────────────────────────────────────────────────────────────────┤ │
│  │ ⬜ react-test-writer                             [Edit] [Delete] │ │
│  │    Generates unit tests for React components                    │ │
│  │    Status: Draft  |  Tags: react, testing  |  v0.1.0            │ │
│  ├─────────────────────────────────────────────────────────────────┤ │
│  │ ⬜ api-docs-generator                            [Edit] [Delete] │ │
│  │    Creates OpenAPI documentation from code                      │ │
│  │    Status: Complete  |  Tags: api, docs  |  v2.1.0              │ │
│  └─────────────────────────────────────────────────────────────────┘ │
│                                                                       │
│  ◀ Previous                                         Page 1/3  Next ▶ │
│                                                                       │
└──────────────────────────────────────────────────────────────────────┘
```

#### 7.5.2 Skill Creation / Conversation View

```
┌──────────────────────────────────────────────────────────────────────┐
│  ████  NEW SKILL  ████                           [Save Draft] [Back] │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │
├─────────────────────────────────┬────────────────────────────────────┤
│                                 │                                    │
│  CONVERSATION                   │  PREVIEW                           │
│  ─────────────────────────────  │  ────────────────────────────────  │
│                                 │                                    │
│  ┌─────────────────────────┐   │  ┌────────────────────────────┐    │
│  │ 🧑 User                  │   │  │ SKILL.md                   │    │
│  │ I need a skill that     │   │  │ ────────────────────────── │    │
│  │ helps write SQL queries │   │  │ ---                        │    │
│  │ for PostgreSQL...       │   │  │ name: sql-query-helper     │    │
│  └─────────────────────────┘   │  │ description: Helps write   │    │
│                                 │  │   and optimize SQL queries │    │
│  ┌─────────────────────────┐   │  │ allowed-tools: Read,Bash   │    │
│  │ 🤖 Assistant             │   │  │ ---                        │    │
│  │ Great! I have a few     │   │  │                            │    │
│  │ clarifying questions:   │   │  │ # SQL Query Helper         │    │
│  │                         │   │  │                            │    │
│  │ 1. What types of        │   │  │ ## Overview                │    │
│  │    queries do you...    │   │  │ This skill assists with... │    │
│  └─────────────────────────┘   │  │                            │    │
│                                 │  │ ## Guidelines              │    │
│  ┌─────────────────────────┐   │  │ - Always use parameterized │    │
│  │ 🧑 User                  │   │  │   queries...               │    │
│  │ Mainly SELECT queries   │   │  │                            │    │
│  │ with complex JOINs...   │   │  └────────────────────────────┘    │
│  └─────────────────────────┘   │                                    │
│                                 │  ┌────────────────────────────┐    │
│  ┌─────────────────────────┐   │  │ Files                      │    │
│  │ Type your message...    │   │  │ ────────────────────────── │    │
│  │                    [➤]  │   │  │ 📄 SKILL.md                │    │
│  └─────────────────────────┘   │  │ 📁 references/             │    │
│                                 │  │    📄 query-patterns.md   │    │
│                                 │  │    📄 optimization.md     │    │
│                                 │  └────────────────────────────┘    │
│                                 │                                    │
│                                 │  [Generate Skill] [Edit Manually]  │
│                                 │                                    │
└─────────────────────────────────┴────────────────────────────────────┘
```

#### 7.5.3 Skill Editor View

```
┌──────────────────────────────────────────────────────────────────────┐
│  ████  EDIT SKILL  ████                          [Save] [Export] [←] │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │
├──────────────────────────────────────────────────────────────────────┤
│  ┌──────────┬────────────┬───────────┬──────────┐                    │
│  │ Overview │ SKILL.md   │ References│ Assets   │                    │
│  └──────────┴────────────┴───────────┴──────────┘                    │
│                                                                       │
│  ┌─────────────────────────────────────────────────────────────────┐ │
│  │ Name:        [sql-query-helper_________________]                │ │
│  │                                                                 │ │
│  │ Description: [Helps write and optimize SQL queries for_______] │ │
│  │              [PostgreSQL databases____________________________] │ │
│  │                                                                 │ │
│  │ Version:     [1.0.0____]    Status: [Complete ▼]               │ │
│  │                                                                 │ │
│  │ Allowed Tools: [Read, Grep, Bash_____________________________] │ │
│  │                                                                 │ │
│  │ Tags:        [sql, database, postgresql______________________] │ │
│  │                                                                 │ │
│  │ Author:      [Your Name_____]   License: [MIT ▼]               │ │
│  └─────────────────────────────────────────────────────────────────┘ │
│                                                                       │
│  ┌─────────────────────────────────────────────────────────────────┐ │
│  │ Conversation History                              [Continue ▶]  │ │
│  │ 12 messages • Last updated: Dec 15, 2024                       │ │
│  └─────────────────────────────────────────────────────────────────┘ │
│                                                                       │
└──────────────────────────────────────────────────────────────────────┘
```

### 7.6 Responsive Considerations

- **Desktop (1200px+)**: Full two-panel layout for conversation view
- **Tablet (768px-1199px)**: Collapsible panels, tabs for conversation/preview
- **Mobile (< 768px)**: Stacked layout, slide-out panels

---

## 8. LLM Integration

### 8.1 System Prompt Design

The LLM interaction uses a carefully designed system prompt to ensure consistent, high-quality skill generation.

```typescript
// src/server/prompts/skillGeneration.ts

export const SKILL_GENERATION_SYSTEM_PROMPT = `
You are an expert Claude Code skill author. Your role is to help users create comprehensive, well-structured Claude Code skills through iterative conversation.

## Your Capabilities
- Generate complete SKILL.md files with proper YAML frontmatter
- Create supporting reference documents for complex skills
- Suggest appropriate assets and templates
- Validate skill structure against Claude Code specifications

## Skill Structure Requirements

Every skill must include a SKILL.md with:
1. YAML Frontmatter (required):
   - name: lowercase letters, numbers, hyphens only (max 64 chars)
   - description: brief explanation (max 1024 chars)

2. Optional Frontmatter:
   - allowed-tools: comma-separated tool restrictions
   - model: target Claude model
   - version: semantic version
   - license: license identifier

3. Markdown Content:
   - Clear overview of skill purpose
   - Step-by-step guidelines
   - Examples of expected inputs/outputs
   - Edge cases and error handling

## Conversation Guidelines

1. **Clarify Before Generating**: Always ask clarifying questions before generating a skill. Key questions include:
   - What specific problem does this skill solve?
   - What tools should Claude have access to?
   - Are there existing patterns or documentation to reference?
   - What should the skill NOT do?

2. **Iterative Refinement**: After generating a preview, ask if the user wants to:
   - Modify the skill description
   - Add or remove sections
   - Include supporting references
   - Adjust tool restrictions

3. **Quality Standards**: Every skill you generate must:
   - Be focused on a single capability
   - Include practical examples
   - Avoid over-complicating simple tasks
   - Follow progressive disclosure (lean SKILL.md, detailed references)

## Response Format

When generating or updating a skill, structure your response as:

1. Brief explanation of changes/additions
2. Full SKILL.md content in a code block
3. Any supporting files (references, assets) in separate code blocks
4. Follow-up questions or suggestions for improvement

## Example Exchange

User: "I need a skill for writing API documentation"
Assistant: "I'd be happy to help you create an API documentation skill! Let me ask a few questions first:

1. What API format are you documenting? (REST, GraphQL, gRPC?)
2. Do you want OpenAPI/Swagger output?
3. Should the skill read existing code to generate docs?"

User: "REST APIs, yes OpenAPI output, and yes read from code"
\`\`\`

### 8.2 Conversation State Management

```typescript
// src/server/services/conversationService.ts

interface ConversationContext {
  conversationId: string;
  skillDraft: SkillDraft | null;
  clarificationPhase: boolean;
  iterationCount: number;
}

interface SkillDraft {
  skillMd: string;
  references: Reference[];
  assets: Asset[];
  scripts: Script[];
}
```

**State Machine:**

```
                     ┌──────────────┐
                     │    START     │
                     └──────┬───────┘
                            │
                            ▼
                     ┌──────────────┐
                     │ CLARIFYING   │◀───────┐
                     │              │        │
                     └──────┬───────┘        │
                            │                │
              User answers  │                │ Need more info
                            ▼                │
                     ┌──────────────┐        │
                     │  GENERATING  │────────┘
                     │              │
                     └──────┬───────┘
                            │
                            ▼
                     ┌──────────────┐
                     │   PREVIEW    │◀───────┐
                     │              │        │
                     └──────┬───────┘        │
                            │                │
              ┌─────────────┼────────────┐   │
              │             │            │   │
         Approve       Modify        Add more
              │             │            │   │
              ▼             └────────────┘   │
       ┌──────────────┐                      │
       │   COMPLETE   │──────────────────────┘
       │              │     Continue refining
       └──────────────┘
```

### 8.3 LLM Request Structure

```typescript
// src/server/services/llmService.ts

interface LLMRequest {
  model: 'claude-sonnet-4-20250514' | 'claude-opus-4-20250514';
  max_tokens: number;
  system: string;
  messages: Message[];
}

interface Message {
  role: 'user' | 'assistant';
  content: string;
}

async function generateSkillResponse(
  context: ConversationContext,
  userMessage: string
): Promise<LLMResponse> {
  const messages = await buildMessageHistory(context.conversationId);
  messages.push({ role: 'user', content: userMessage });

  const response = await anthropic.messages.create({
    model: 'claude-sonnet-4-20250514',
    max_tokens: 4096,
    system: SKILL_GENERATION_SYSTEM_PROMPT,
    messages: messages
  });

  return parseResponse(response);
}
```

### 8.4 Response Parsing

The LLM response parser extracts structured skill content from the assistant's response:

```typescript
interface ParsedResponse {
  textContent: string;
  skillMd: string | null;
  references: ParsedFile[];
  assets: ParsedFile[];
  scripts: ParsedFile[];
  isQuestion: boolean;
}

function parseResponse(response: LLMResponse): ParsedResponse {
  // Extract code blocks with skill content
  // Identify questions vs statements
  // Parse file paths and content
}
```

### 8.5 Error Handling

| Error Type | Handling Strategy |
|------------|-------------------|
| API Rate Limit | Exponential backoff with user notification |
| Invalid Response | Request regeneration with clarification |
| Context Too Long | Summarize earlier messages, preserve key decisions |
| Network Error | Retry with timeout, offline message queuing |

---

## 9. Skill Export Format

### 9.1 Archive Structure

The `.skill` export format is a ZIP archive with the following structure:

```
skill-name.skill
├── SKILL.md              # Main skill definition
├── manifest.json         # Metadata and file manifest
├── references/           # Documentation files
│   ├── schemas.md
│   └── patterns.md
├── assets/               # Binary and template files
│   ├── template.html
│   └── icon.png
└── scripts/              # Executable scripts
    ├── process.py
    └── helper.sh
```

### 9.2 Manifest Schema

```json
{
  "$schema": "https://skills-manager.local/manifest.schema.json",
  "version": "1.0",
  "skill": {
    "name": "sql-query-helper",
    "description": "Helps write and optimize SQL queries",
    "version": "1.0.0",
    "author": "Developer Name",
    "license": "MIT",
    "created_at": "2024-12-15T10:30:00Z",
    "updated_at": "2024-12-15T14:22:00Z"
  },
  "files": {
    "skill_md": "SKILL.md",
    "references": [
      {
        "path": "references/schemas.md",
        "description": "Database schema documentation"
      },
      {
        "path": "references/patterns.md",
        "description": "Common query patterns"
      }
    ],
    "assets": [
      {
        "path": "assets/template.html",
        "content_type": "text/html",
        "size_bytes": 2048
      }
    ],
    "scripts": [
      {
        "path": "scripts/process.py",
        "language": "python",
        "description": "Data processing utility"
      }
    ]
  },
  "checksum": "sha256:abc123..."
}
```

### 9.3 Export Service Implementation

```typescript
// src/server/services/exportService.ts

import archiver from 'archiver';

async function exportSkill(skillId: string): Promise<Buffer> {
  const skill = await skillRepository.getById(skillId);
  const references = await referenceRepository.getBySkillId(skillId);
  const assets = await assetRepository.getBySkillId(skillId);
  const scripts = await scriptRepository.getBySkillId(skillId);

  const archive = archiver('zip', { zlib: { level: 9 } });
  const chunks: Buffer[] = [];

  archive.on('data', (chunk) => chunks.push(chunk));

  // Add SKILL.md
  archive.append(skill.skill_md_content, { name: 'SKILL.md' });

  // Add references
  for (const ref of references) {
    archive.append(ref.content, { name: ref.file_path });
  }

  // Add assets
  for (const asset of assets) {
    archive.append(asset.content, { name: asset.file_path });
  }

  // Add scripts
  for (const script of scripts) {
    archive.append(script.content, { name: script.file_path });
  }

  // Add manifest
  const manifest = buildManifest(skill, references, assets, scripts);
  archive.append(JSON.stringify(manifest, null, 2), { name: 'manifest.json' });

  await archive.finalize();

  return Buffer.concat(chunks);
}
```

### 9.4 Import Service Implementation

```typescript
async function importSkill(archiveBuffer: Buffer): Promise<Skill> {
  const zip = await JSZip.loadAsync(archiveBuffer);

  // Validate manifest
  const manifestFile = zip.file('manifest.json');
  if (!manifestFile) throw new Error('Invalid .skill archive: missing manifest');

  const manifest = JSON.parse(await manifestFile.async('string'));
  validateManifest(manifest);

  // Extract and validate SKILL.md
  const skillMdFile = zip.file('SKILL.md');
  if (!skillMdFile) throw new Error('Invalid .skill archive: missing SKILL.md');

  const skillMdContent = await skillMdFile.async('string');
  validateSkillMd(skillMdContent);

  // Create skill record
  const skill = await skillRepository.create({
    name: manifest.skill.name,
    description: manifest.skill.description,
    skill_md_content: skillMdContent,
    version: manifest.skill.version,
    author: manifest.skill.author,
    license: manifest.skill.license,
    status: 'complete'
  });

  // Import references, assets, scripts...

  return skill;
}
```

---

## 10. Security Considerations

### 10.1 API Key Management

| Concern | Mitigation |
|---------|------------|
| Key Storage | Store encrypted in settings table using system keychain or environment variable |
| Key Exposure | Never log or display full API key; show only last 4 characters |
| Key Rotation | Provide easy mechanism to update key without data loss |

```typescript
// Settings encryption approach
import { createCipheriv, createDecipheriv, randomBytes } from 'crypto';

const ENCRYPTION_KEY = process.env.ENCRYPTION_KEY || generateMachineKey();
const ALGORITHM = 'aes-256-gcm';

function encryptApiKey(apiKey: string): string {
  const iv = randomBytes(16);
  const cipher = createCipheriv(ALGORITHM, ENCRYPTION_KEY, iv);
  const encrypted = Buffer.concat([cipher.update(apiKey), cipher.final()]);
  const authTag = cipher.getAuthTag();
  return `${iv.toString('hex')}:${authTag.toString('hex')}:${encrypted.toString('hex')}`;
}
```

### 10.2 Input Validation

All user inputs must be validated:

```typescript
// src/shared/validators/skillValidator.ts

import { z } from 'zod';

const skillNameSchema = z
  .string()
  .min(1)
  .max(64)
  .regex(/^[a-z0-9-]+$/, 'Name must contain only lowercase letters, numbers, and hyphens');

const skillDescriptionSchema = z
  .string()
  .min(10)
  .max(1024);

const skillSchema = z.object({
  name: skillNameSchema,
  description: skillDescriptionSchema,
  skill_md_content: z.string().min(50),
  allowed_tools: z.string().optional(),
  tags: z.string().optional(),
  version: z.string().regex(/^\d+\.\d+\.\d+$/).optional()
});
```

### 10.3 File Upload Security

For asset uploads:

| Check | Implementation |
|-------|----------------|
| File Size | Limit to 10MB per file, 50MB total per skill |
| File Type | Whitelist allowed MIME types |
| Filename | Sanitize to prevent path traversal |
| Content | Scan for malicious content patterns |

```typescript
const ALLOWED_MIME_TYPES = [
  'text/plain',
  'text/markdown',
  'text/html',
  'text/css',
  'application/json',
  'application/javascript',
  'image/png',
  'image/jpeg',
  'image/svg+xml'
];

const MAX_FILE_SIZE = 10 * 1024 * 1024; // 10MB

function validateAssetUpload(file: UploadedFile): void {
  if (file.size > MAX_FILE_SIZE) {
    throw new Error(`File exceeds maximum size of ${MAX_FILE_SIZE} bytes`);
  }

  if (!ALLOWED_MIME_TYPES.includes(file.mimetype)) {
    throw new Error(`File type ${file.mimetype} is not allowed`);
  }

  // Sanitize filename
  const sanitizedName = path.basename(file.name).replace(/[^a-zA-Z0-9._-]/g, '_');
  file.name = sanitizedName;
}
```

### 10.4 Local-Only Access

The application runs locally and should not be exposed to the network:

```typescript
// src/server/index.ts

const app = express();

// Bind only to localhost
const HOST = '127.0.0.1';
const PORT = process.env.PORT || 3000;

app.listen(PORT, HOST, () => {
  console.log(`Skills Manager running at http://${HOST}:${PORT}`);
});
```

### 10.5 Content Security Policy

```typescript
// src/server/middleware/security.ts

import helmet from 'helmet';

app.use(helmet({
  contentSecurityPolicy: {
    directives: {
      defaultSrc: ["'self'"],
      scriptSrc: ["'self'"],
      styleSrc: ["'self'", "'unsafe-inline'"], // Required for Tailwind
      imgSrc: ["'self'", "data:"],
      connectSrc: ["'self'", "https://api.anthropic.com"],
      fontSrc: ["'self'"],
      objectSrc: ["'none'"],
      frameAncestors: ["'none'"]
    }
  }
}));
```

---

## 11. Implementation Phases

### Phase 1: Foundation (Core Infrastructure)

**Goal:** Establish basic project structure and database connectivity

**Tasks:**
1. Initialize Node.js project with TypeScript configuration
2. Set up Vite for frontend bundling
3. Configure TailwindCSS with custom Minecraft theme
4. Implement DuckDB connection and basic repository pattern
5. Create database migrations for all tables
6. Build basic Express server with error handling
7. Implement settings API for API key storage

**Deliverables:**
- Working development environment
- Database schema deployed
- Basic API structure
- Settings persistence

**Success Criteria:**
- Application starts without errors
- Database operations work correctly
- API key can be saved and retrieved

---

### Phase 2: Skill CRUD (Basic Management)

**Goal:** Implement core skill management without LLM

**Tasks:**
1. Build Skills API endpoints (list, get, create, update, delete)
2. Implement skill validation logic
3. Create SkillList frontend component with Minecraft styling
4. Create SkillEditor frontend component
5. Implement search and filter functionality
6. Add pagination to skill list
7. Implement skill duplication

**Deliverables:**
- Functional skill CRUD operations
- Styled skill list view
- Skill editor with form validation

**Success Criteria:**
- Can create, edit, delete skills manually
- Search finds skills by name/description
- UI matches Minecraft theme

---

### Phase 3: LLM Integration (AI-Powered Creation)

**Goal:** Integrate Claude API for skill generation

**Tasks:**
1. Implement LLMService with Claude API client
2. Create system prompt for skill generation
3. Build ConversationService for state management
4. Implement Conversations API endpoints
5. Create ConversationView frontend component
6. Implement message streaming (optional)
7. Build response parser for skill extraction
8. Add preview panel for generated skills

**Deliverables:**
- Working conversation flow
- AI-generated SKILL.md content
- Iterative refinement capability

**Success Criteria:**
- Can describe skill and receive generated content
- AI asks clarifying questions
- Generated skills are valid

---

### Phase 4: References & Assets

**Goal:** Support complete skill packages

**Tasks:**
1. Implement References API and repository
2. Implement Assets API with file upload
3. Implement Scripts API and repository
4. Extend LLM prompt to suggest supporting files
5. Add reference editor tab in SkillEditor
6. Add asset manager tab with upload/download
7. Add scripts editor tab
8. Update conversation flow to include file suggestions

**Deliverables:**
- Full support for skill directory structure
- File upload/download functionality
- LLM-suggested supporting files

**Success Criteria:**
- Can add/edit/delete references
- Can upload and manage assets
- LLM suggests appropriate supporting files

---

### Phase 5: Export & Import

**Goal:** Enable skill deployment

**Tasks:**
1. Implement ExportService for single skill
2. Implement batch export functionality
3. Create .skill archive format with manifest
4. Implement ImportService with validation
5. Add export button to skill list
6. Add export modal for batch selection
7. Add import functionality to main view

**Deliverables:**
- Working .skill export format
- Single and batch export
- Import with validation

**Success Criteria:**
- Exported skills are valid ZIP archives
- Imported skills appear in list
- Round-trip export/import preserves all data

---

### Phase 6: Polish & UX

**Goal:** Refine user experience

**Tasks:**
1. Add loading states and progress indicators
2. Implement error messages and toast notifications
3. Add keyboard shortcuts
4. Improve responsive design
5. Add help/documentation page
6. Performance optimization
7. Add conversation history continuation
8. Implement skill templates/presets

**Deliverables:**
- Polished user interface
- Comprehensive error handling
- Documentation

**Success Criteria:**
- No unhandled errors in UI
- Responsive on tablet/mobile
- Users can self-serve with documentation

---

## 12. Appendices

### Appendix A: Glossary

| Term | Definition |
|------|------------|
| Skill | A Claude Code extension that provides specialized knowledge or capabilities |
| SKILL.md | The main markdown file defining a skill's purpose and instructions |
| Frontmatter | YAML metadata at the top of SKILL.md |
| Progressive Disclosure | Information architecture pattern where detail is revealed as needed |
| DuckDB | Embedded analytical database engine |
| Claude API | Anthropic's API for accessing Claude models |

### Appendix B: Technology References

- **DuckDB Node.js**: https://duckdb.org/docs/api/nodejs/overview
- **TailwindCSS**: https://tailwindcss.com/docs
- **Claude API**: https://docs.anthropic.com/claude/reference
- **Vite**: https://vitejs.dev/guide/
- **Express**: https://expressjs.com/
- **Archiver**: https://www.archiverjs.com/

### Appendix C: Minecraft CSS Resources

- **Minecraft-CSS Framework**: https://github.com/Jiyath5516F/Minecraft-CSS
- **VT323 Font (Pixel style)**: https://fonts.google.com/specimen/VT323
- **Press Start 2P Font**: https://fonts.google.com/specimen/Press+Start+2P
- **Minecraft Color Palette**: Various community resources for authentic colors

### Appendix D: Sample SKILL.md

```markdown
---
name: react-component-generator
description: Generates React functional components with TypeScript, including tests and stories
allowed-tools: Read, Write, Glob
version: 1.0.0
---

# React Component Generator

This skill helps you quickly scaffold React components with all the necessary files.

## What It Creates

For each component, this skill generates:
1. Component file with TypeScript types
2. CSS module for styling
3. Unit test file with React Testing Library
4. Storybook story file

## Usage

Simply describe the component you need:

**Example**: "Create a Button component that supports primary, secondary, and danger variants with an optional loading state"

## Guidelines

- Components are created in `src/components/` by default
- Uses functional components with hooks
- Follows project naming conventions
- Includes accessibility attributes
- Types are co-located with components

## Limitations

- Does not modify existing components
- Does not handle state management setup
- Assumes standard React project structure
```

### Appendix E: Environment Variables

```bash
# .env.example

# Server Configuration
PORT=3000
HOST=127.0.0.1

# Database
DB_PATH=./data/skills.db

# Encryption (generate with: openssl rand -hex 32)
ENCRYPTION_KEY=your-32-byte-hex-key

# Claude API (alternatively stored in settings)
ANTHROPIC_API_KEY=sk-ant-...

# Development
NODE_ENV=development
```

---

## Document History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0.0 | December 2024 | Claude | Initial specification |

---

*End of Specification*
