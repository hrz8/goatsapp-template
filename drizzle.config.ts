
import dotenv from 'dotenv';
dotenv.config();

import type {Config} from 'drizzle-kit';

const DATABASE_URL = process.env.DATABASE_URL;
const MIGRATION_FOLDER = './db/drizzle';

if (!DATABASE_URL) {
  throw new Error('cannot resolve database connection!');
}

export default {
  schema: './db/**/*.schema.ts',
  out: MIGRATION_FOLDER,
  driver: 'pg',
  dbCredentials: {
    connectionString: DATABASE_URL,
  },
} satisfies Config;
