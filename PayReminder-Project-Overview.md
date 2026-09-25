# PayReminder — Project Documentation

## What This Project Is

PayReminder is a lightweight SaaS tool built for freelancers and small agencies who struggle to get paid on time. Instead of manually tracking invoices and awkwardly following up with clients, PayReminder automates the entire follow-up process — from a friendly first nudge to a firm final notice — and stops the moment payment comes in.

This is being built as a solo, real-world SaaS product (not just a practice project) — the goal is to eventually have paying users in the UK/US/global freelance market, while also serving as a strong portfolio piece.

---

## The Problem

Freelancers and small agencies typically:
- Track invoices manually in spreadsheets
- Forget to follow up, letting payments slip weeks past due
- Feel awkward directly asking clients for money they're already owed
- Suffer real cash flow problems as a result

Existing tools are either too basic (a plain invoice generator) or too bloated (full accounting suites like QuickBooks) for this one specific pain point.

---

## The Solution

A focused tool that does one thing well: **get freelancers paid on time, without the awkward follow-up.**

Users add a client and an invoice once. From there, PayReminder handles the rest automatically.

---

## Target Audience

- Solo freelancers — designers, developers, writers, consultants, marketers
- Small agencies (2–10 people)
- Primarily based in or billing clients in the UK, US, Canada, Australia, and EU

---

## How It Works (User Flow)

1. User signs up and sets up their business profile
2. Adds a client (name, email)
3. Creates an invoice (amount, currency, due date, description)
4. Marks the invoice as "Sent" — the reminder schedule kicks in automatically
5. The system sends reminders on a fixed schedule, growing firmer over time:
   - 3 days before due → friendly reminder
   - On the due date → "due today" reminder
   - 3–7 days overdue → firmer follow-up
   - 14 days overdue → final notice
6. Every reminder includes a one-click Stripe payment link
7. When the client pays, the invoice is marked "Paid" automatically and reminders stop
8. The user sees everything at a glance on their dashboard — what's outstanding, what's overdue, what's been paid

---

## Core Features

### Must-have (MVP)
- **Account & business profile** — sign up, log in, basic business details
- **Client management** — add, edit, remove clients
- **Invoice management** — create invoices with amount, currency, due date, and description
- **Invoice status tracking** — Due → Paid / Cancelled, with rules so a paid or cancelled invoice can't be changed again
- **Automated reminders** — escalating email sequence, fully automatic once an invoice is marked sent
- **One-click payment** — Stripe payment link embedded in every reminder email
- **Auto-stop on payment** — reminders stop the instant payment is confirmed
- **Dashboard** — at-a-glance view of total outstanding, overdue, and paid amounts

### Planned for later (Phase 2)
- Custom reminder schedules (user sets their own timing/tone)
- Multi-currency support
- Late fee auto-calculation
- Custom email templates
- Recurring invoices

### Longer-term (Phase 3)
- Quotes/proposals (before an invoice even exists)
- Team accounts with roles (for agencies)
- White-label branding
- Third-party integrations (e.g. Zapier)
- Reporting and analytics

---

## Business Model

| Tier | Price | What's included |
|---|---|---|
| Free | $0/month | Up to 3 active invoices |
| Pro | $15/month | Unlimited invoices, full reminder automation |
| Annual Pro | ~$12/month (billed yearly) | Same as Pro, discounted |

---

## What Success Looks Like

- A working MVP a real freelancer could use to replace their spreadsheet-and-memory system
- At least a handful of real signups/paying users validating the idea
- A polished, demonstrable project for freelance/job portfolio purposes

---

## Current Status

🚧 In active development — building step by step, starting with core invoice and client management.
