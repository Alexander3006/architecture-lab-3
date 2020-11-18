CREATE TABLE "Machines"(
                           "id" SERIAL PRIMARY KEY,
                           "name" VARCHAR NOT NULL,
                           "cpuCount" smallint check ( "cpuCount" > 0 ) NOT NULL
);

CREATE TABLE "Disks" (
                         "id" SERIAL PRIMARY KEY,
                         "space" INTEGER CHECK ("space" > 0) NOT NULL,
                         "machineId" INTEGER DEFAULT NULL,
                         FOREIGN KEY ("machineId") REFERENCES "Machines"("id") ON DELETE SET NULL
);

CREATE OR REPLACE FUNCTION "MachinesInfo"()
    returns TABLE (
                      id int,
                      "name" varchar,
                      "cpuCount" smallint,
                      "totalDiskSpace" bigint
                  )
    language plpgsql
as $code$
BEGIN
    RETURN QUERY SELECT M.id, M.name, M."cpuCount", sum(D.space) as "totalDiskSpace"
                 FROM "Machines" as M
                          LEFT JOIN "Disks" as D on M.id = D."machineId"
                 GROUP BY M.id;
END
$code$;

