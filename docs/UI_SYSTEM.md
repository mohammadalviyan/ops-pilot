# UI System — OpsPilot

## 1. Design Direction

OpsPilot uses **Ant Design** with a subtle **OpsPilot Glass Style**.

The interface should feel like:

- RPA operations command center.
- Internal monitoring cockpit.
- Professional support dashboard.
- Clean but not empty.
- Modern but not decorative.
- Data-first and readable.
- Useful for real support, escalation, and reporting work.

The UI should help users answer:

- What needs attention today?
- Which robots failed?
- Which tickets are waiting?
- Which license or asset is risky?
- What should the support team do next?

---

## 2. UI Personality

OpsPilot should feel:

- Professional.
- Calm.
- Operational.
- Helpful.
- Structured.
- Slightly modern/tech.
- Human-designed, not AI-template generated.

Avoid UI that feels:

- Generic SaaS template.
- Plain default Ant Design without identity.
- Too colorful.
- Too decorative.
- Too empty.
- Too many identical cards.
- Too much glass effect.
- Too much white space without information hierarchy.

---

## 3. Ant Design Usage

Ant Design is the primary component library.

Prefer Ant Design for:

- Layout.
- Menu.
- Card.
- Button.
- Form.
- Input.
- Select.
- DatePicker.
- Table.
- Tag.
- Badge.
- Dropdown.
- Modal.
- Drawer.
- Tabs.
- Timeline.
- Statistic.
- Progress.
- Alert.
- Empty.
- Skeleton.
- App message / modal / notification.

Do not build custom versions of these components unless the active task explicitly requires it.

Do not mix another UI component library into the frontend.

---

## 4. Glass Style Usage

Glass effect should be used subtly and intentionally.

Good places for glass:

- Sidebar.
- Topbar.
- Dashboard stat cards.
- Summary cards.
- Filter panels.
- Empty state containers.
- Important page-level panels.

Avoid glass effect for:

- Every table cell.
- Every button.
- Every form field.
- Long dense forms.
- Long text areas.
- Error/critical messages.
- Content that needs maximum contrast.

Glass style must improve depth and visual hierarchy, not reduce readability.

---

## 5. Visual Tokens

Use these directions through `frontend/config/antd-theme.ts`.

### Color Direction

| Purpose | Direction |
|---|---|
| Primary | Calm blue |
| Success | Green |
| Warning | Amber |
| Danger | Red |
| Neutral | Slate / gray |
| Background | Soft light blue-gray |
| Border | Subtle slate transparency |

### Shape

- Cards: rounded but not playful.
- Buttons: medium rounded.
- Tags: pill-like where appropriate.
- Tables: readable row height.
- Modals/drawers: clean and focused.

### Shadow

Use soft shadow only for depth. Avoid heavy floating-card effects.

### Background

Use subtle radial gradients sparingly for the app background. Avoid strong gradient blocks.

---

## 6. Required Base Components

The frontend should establish these reusable components before building complex feature pages.

```txt
components/layout/
  AppShell.tsx
  Sidebar.tsx
  Topbar.tsx
  PageContainer.tsx
  PageHeader.tsx

components/glass/
  GlassCard.tsx
  GlassPanel.tsx
  GlassStatCard.tsx
  GlassPageHeader.tsx

components/data-display/
  StatusTag.tsx
  PriorityTag.tsx
  DataTable.tsx
  FilterToolbar.tsx
  ActivityTimeline.tsx
  MetricGroup.tsx

components/feedback/
  EmptyState.tsx
  LoadingState.tsx
  ErrorState.tsx

components/forms/
  FormSection.tsx
  FieldError.tsx

components/charts/
  ChartCard.tsx
  TicketStatusChart.tsx
  RobotRunChart.tsx
  TicketTrendChart.tsx
```

---

## 7. Page Layout Rules

Every page should include:

- Clear page title.
- Short helper text under the title.
- Primary action on the top-right if needed.
- Main content area.
- Loading state.
- Error state.
- Empty state when applicable.
- Consistent spacing.
- Consistent status/priority tags.
- Desktop and tablet-friendly layout.

Page header example:

```txt
Tickets
Track RPA issues, requests, escalations, and operational follow-ups.
[Create Ticket]
```

Avoid page headers like:

```txt
Manage Data
Lorem ipsum dolor sit amet
```

---

## 8. Dashboard Rules

Dashboard cards should not all look identical.

Each stat card should include:

- Label.
- Value.
- Short context.
- Optional trend.
- Optional severity/status indicator.

Good:

```txt
Open Tickets
24
7 waiting for user response
```

Bad:

```txt
Open Tickets
24
```

Dashboard should prioritize:

1. Tickets needing attention.
2. Failed robot runs.
3. SLA risk.
4. Expiring licenses/assets.
5. Monthly trends.

---

## 9. Status and Priority Rules

Status and priority must be shown consistently through wrapper components.

Use `StatusTag` for statuses and `PriorityTag` for priorities.

### Ticket Status

| Status | Visual Direction |
|---|---|
| OPEN | Neutral |
| IN_PROGRESS | Blue / active |
| WAITING_USER | Amber / waiting |
| WAITING_INTERNAL | Amber / waiting |
| WAITING_VENDOR | Amber / waiting |
| RESOLVED | Green / success |
| CLOSED | Gray / muted |
| CANCELLED | Gray / muted |

### Robot Run Status

| Status | Visual Direction |
|---|---|
| RUNNING | Blue / active |
| SUCCESS | Green / success |
| FAILED | Red / danger |
| PARTIAL_SUCCESS | Amber / warning |
| CANCELLED | Gray / muted |
| SKIPPED | Gray / muted |

### Priority

| Priority | Visual Direction |
|---|---|
| LOW | Neutral |
| MEDIUM | Blue / normal |
| HIGH | Amber / important |
| CRITICAL | Red / urgent |

Do not manually style tags per page. Always reuse the shared components.

---

## 10. Table Rules

Tables should include:

- Search/filter area.
- Clear column names.
- Status tag.
- Priority tag when applicable.
- Row action.
- Pagination.
- Loading state.
- Empty state.
- Truncated long text.
- Readable row height.

Avoid:

- Too many columns at once.
- Long text overflowing.
- Icon-only actions without tooltip or label.
- Repeated custom table implementations.

Use Ant Design `Table` as the base.

---

## 11. Form Rules

Forms should be grouped by meaning.

Recommended grouping:

- Basic Information.
- Classification.
- Assignment.
- Related Object.
- Timeline / SLA.
- Notes.

Use helper text for fields that are not obvious.

Every form should handle:

- Validation.
- Submit loading state.
- Error feedback.
- Success feedback.
- Cancel/back action.

Use Ant Design `Form` as the base.

---

## 12. Timeline Rules

Timeline should be used for ticket activities, comments, and escalations.

Timeline items should show:

- Action type.
- Actor.
- Timestamp.
- Short description.
- Internal/public indicator when relevant.

Good copy:

```txt
Status changed from OPEN to IN_PROGRESS by RPA Support.
```

Avoid:

```txt
Data updated successfully.
```

---

## 13. Empty State Copy

Empty states should sound human and operational.

Good examples:

```txt
No tickets need attention right now.
No robot run logs recorded yet.
No assets are expiring soon. Everything looks safe for now.
No comments yet. Add the first update to keep the history clear.
```

Avoid:

```txt
Data not found.
No data.
Empty.
```

---

## 14. Loading and Error States

Use `LoadingState` and `ErrorState` consistently.

Loading should be calm and not disruptive.

Error messages should be specific when possible.

Good:

```txt
Unable to load tickets. Please try again.
```

Bad:

```txt
An error occurred.
```

---

## 15. Copywriting Tone

Use clear operational language.

Prefer:

- Waiting for user response.
- Failed today.
- Expiring soon.
- Needs attention.
- No critical issue detected.
- Track RPA issues and follow-ups.
- Review failed runs and linked tickets.

Avoid:

- Manage your data.
- Submit form.
- Data processed.
- Error occurred.
- Lorem ipsum.

---

## 16. Frontend Agent Checklist

Before implementing a visible frontend task, the agent must check:

- Is there an existing reusable component?
- Does the page need loading, empty, and error states?
- Does the UI follow OpsPilot Glass Style?
- Are status and priority displayed using shared components?
- Is the page useful for real RPA support work?
- Is the copy human and operational?
- Does the implementation avoid unnecessary custom components?

After implementing, the agent must update:

- Active task file.
- `backlog/STATUS.md`.
- `docs/UI_SYSTEM.md` if reusable UI rules or components changed.
- `docs/DECISIONS.md` if a frontend technical decision changed.
