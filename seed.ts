import dotenv from 'dotenv';
dotenv.config();

import { drizzle } from 'drizzle-orm/node-postgres';
import { Pool } from 'pg';
import { projects } from './db/schemas/projects.schema';

const DATABASE_URL = process.env.DATABASE_URL;
if (!DATABASE_URL) {
  throw new Error('cannot resolve database connection!');
}
void (async function (): Promise<void> {
  const client = new Pool({
    connectionString: DATABASE_URL,
  });
  const db = drizzle(client);
  const data: (typeof projects.$inferInsert)[] = [];

  data.push({
    id: 'de640718-efc8-4eef-9d96-3dea2e3fe827',
    name: 'Example Project',
    alias: 'example-project',
    description: 'This is only an example project to start with.',
    settings: {
      incomingURL: 'https://your.website.com/whatsapp/web-hook'
    }
  });

  await db.insert(projects).values(data);
})();