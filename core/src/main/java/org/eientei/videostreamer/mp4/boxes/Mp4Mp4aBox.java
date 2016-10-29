package org.eientei.videostreamer.mp4.boxes;

import io.netty.buffer.ByteBuf;
import org.eientei.videostreamer.mp4.Mp4AudioTrackAac;
import org.eientei.videostreamer.mp4.Mp4Box;
import org.eientei.videostreamer.mp4.Mp4RemuxerHandler;
import org.eientei.videostreamer.mp4.Mp4Track;

import java.util.List;

/**
 * Created by Alexander Tumin on 2016-10-23
 */
public class Mp4Mp4aBox extends Mp4Box {
    private final Mp4AudioTrackAac track;
    private final Mp4EsdsBox esds;

    public Mp4Mp4aBox(Mp4RemuxerHandler context, List<Mp4Track> tracks, Mp4AudioTrackAac track) {
        super("mp4a", context);
        this.esds = new Mp4EsdsBox(context, tracks, track);
        this.track = track;
    }

    @Override
    protected void doWrite(ByteBuf out) {
        out.writeInt(0);
        out.writeShort(0);

        out.writeShort(1);

        out.writeInt(0);
        out.writeInt(0);

        out.writeShort(track.getChannels());
        out.writeShort(track.getSampleSize());

        out.writeInt(0);
        out.writeShort((int)track.getTimescale());
        out.writeShort(track.getSampleRate());
        esds.write(out);
    }
}